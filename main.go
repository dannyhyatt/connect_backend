package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"database/sql"

	"github.com/go-redis/redis"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/sfreiberg/gotwilio"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	HOSTED_URL = "http://71.191.95.144:3000/"
)

var db *sql.DB
var twilio *gotwilio.Twilio
var sg_client *sendgrid.Client
var rd_client *redis.Client

func main() {

	// init is automatically called, thanks golang
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong" + "c",
			})
		})

		api.GET("/usernameAvailable/:username", func(c *gin.Context) {

			// todo if username is taken but account is not verified check if the verification has timed out
			// and if so remove the account

			query := "SELECT * FROM users WHERE username='" + c.Param("username") + "'"
			rows, err := db.Query(query)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H {
					"success" : false,
				})

				fmt.Print("error: ")
				fmt.Println(err)

				return
			}

			if rows.Next() {
				c.JSON(http.StatusOK, gin.H {
					"success" : false,
					"error" : "Username taken",
				})
			} else {
				c.JSON(http.StatusOK, gin.H {
					"success" : true,
				})
			}
		})

		api.POST("/register", func(c *gin.Context) {
			username := c.PostForm("username")
			password := c.PostForm("password") // todo hash passwords
			name := c.PostForm("full_name")
			email := c.PostForm("email")
			phoneNumber := c.PostForm("phone_number")

			query := "INSERT INTO users (username, password_hash, name, email, phone_number) VALUES ('" + username + "', '" + password + "', '" + name + "', '" + email + "', '" + phoneNumber + "')"

			// probably could do the below with Sprintf
			fmt.Printf("query: %s\n\n", query)
			_, err := db.Exec(query)
			fmt.Println("RRPRPRR")

			if err != nil {
				c.JSON(http.StatusNotAcceptable, gin.H {
					"success" : false,
					"error" : "Your credentials are unacceptable", // todo provide a meaningful error & check if the error is server side connecting to the db
				})

				fmt.Print("error: ")
				fmt.Print(err)

				return
			} else {
				//id:=_
				//fmt.Print("id: ")
				//fmt.Println(id)
			}

			// now set verification codes

			phoneCode := strconv.Itoa(rand.Intn(8999) + 1000)
			emailCode := randStringBytesMaskImprSrcSB(24)

			err = rd_client.Set(username + "_phoneCode", []byte(phoneCode), time.Minute * 30).Err()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H {
					"success" : false,
					"error" : "Internal Server Error",
				})
				return
			}

			err = rd_client.Set(username + "_emailCode", []byte(emailCode), time.Hour * 30).Err() // set to "VERIFIED" when verified for both
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H {
					"success" : false,
					"error" : "Internal Server Error",
				})
				return
			}

			from := "+12406410605"
			to   := phoneNumber
			message := "Connect here! Your verification code is " + phoneCode
			// todo catch these errors before inserting into db
			smsResponse, exception, err := twilio.SendSMS(from, to, message, "", "")
			fmt.Println("sms, exception, err:")
			fmt.Println(smsResponse)
			fmt.Println(exception)
			fmt.Println(err)
			if err != nil {
				fmt.Println("ERROR SENDING EMAIL")
				c.JSON(http.StatusOK, gin.H {
					"success" : false,
					"error" : "Error sending text. Try using another phone number?",
				})
				return
			}
			err = sendEmail(email, "Your Connect Registration Code",
				"Please go to the URL here to verify your email:\n http://"+HOSTED_URL+"api/attemptVerifyEmail/"+emailCode,
				"<img src=\"" + HOSTED_URL + "/EMAIL_LOGO.PNG\"><br><h1>Welcome to Connect!<h1> <h6>Click <a href=\"" + HOSTED_URL + "api/attemptVerifyEmail/" + emailCode + "\">here</a> to verify your account.</h6>",
				"Danny", "Danny from Connect", name)
			if err != nil {
				fmt.Println("ERROR SENDING EMAIL")
				c.JSON(http.StatusOK, gin.H {
					"success" : false,
					"error" : "Error sending email. Try using another email address?",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H {
				"success" : true,
			})

		})

		// do we need to send credentials to verify a phone number
		api.GET("/attemptVerifyPhone/:username/:code", func(c *gin.Context) {
			item, err := rd_client.Get(c.Param("username") + "_phoneCode").Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H {
					"success" : false,
					"error" : "Internal server error. Do you have an account?",
				})
				return
			}

			//var created string
			//created, err = rd_client.Get(c.Param("username") + "_phoneCode").Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H {
					"success" : false,
					"error" : "Internal server error. Do you have an account?",
				})
				return
			}

			fmt.Print("attempt: ")
			fmt.Print(item)
			fmt.Print(", ")
			fmt.Println(c.Param("code"))

			if item == c.Param("code") {
				err = rd_client.Set(c.Param("username") + "_phoneCode", []byte("VERIFIED"), time.Minute * 30).Err()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H {
						"success" : false,
						"error" : "Internal server error",
					})
				}
				c.JSON(http.StatusOK, gin.H{
					"success" : true,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"success" : false,
					"error" : "Wrong validation code",
				})
			}
		})

		api.GET("attemptVerifyEmail/:username/:code", func(c *gin.Context) {
			item, err := rd_client.Get(c.Param("username") + "_emailCode").Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H {
					"success" : false,
					"error" : "Internal server error. Do you have an account?",
				})
				return
			}

			//var created string
			//created, err = rd_client.Get(c.Param("username") + "_phoneCode").Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H {
					"success" : false,
					"error" : "Internal server error. Do you have an account?",
				})
				return
			}

			fmt.Print("attempt: ")
			fmt.Print(item)
			fmt.Print(", ")
			fmt.Println(c.Param("code"))

			if item == c.Param("code") {
				err = rd_client.Set(c.Param("username") + "_emailCode", []byte("VERIFIED"), time.Minute * 30).Err()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H {
						"success" : false,
						"error" : "Internal server error",
					})
				}
				c.JSON(http.StatusOK, gin.H{
					"success" : true,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"success" : false,
					"error" : "Wrong validation code",
				})
			}
		})
	}

	// Start and run the server
	print(router.Run("0.0.0.0:3000"))
}

//func registerUser(username string, password string, email string, phoneNumber string) error {
//	if !(len(username) > 3 && len(username) < 15) {
//		return errors.New("username must be between 4 and 14 characters")
//	}
//
//	if !(len(password) < 65 && len(password) >= 8) {
//		return errors.New("password should be between 8 and 64 characters")
//	}
//
//	if !isValidEmail(email) {
//		return errors.New("invalid email address")
//	}
//
//	if !isValidPhoneNumber(phoneNumber) {
//		return errors.New("invalid phone number")
//	}
//
//	//if sendRegistrationEmail(email) != nil {
//	//	return errors.New("error sending confirmation email")
//	//}
//
//	if sendRegistrationText(phoneNumber) != nil {
//		return errors.New("error sending confirmation text")
//	}
//
//	return nil
//}

func isValidEmail(address string) bool {
	// todo
	return strings.LastIndex(address, "@") != -1
}

func isValidPhoneNumber(phoneNumber string) bool {
	// todo
	return true
}

// todo check if email is verified after phone is verified
func phoneVerifiedCheckEmailVerified(username string) {

}

// todo check if phone is verified after email is verified
func emailVerifiedCheckPhoneVerified(username string) {

}

func verifyAccount(username string) {
	// todo
}

func sendEmail(address, subject, textBody, htmlBody, fromPrefix, fromName, toName string) error {
	if fromPrefix == "" {
		fromPrefix = "connect"
	}
	from := mail.NewEmail(fromName, fromPrefix + "@connect.charity")
	to := mail.NewEmail(toName, address)
	if htmlBody == "" {
		htmlBody = textBody
	}
	message := mail.NewSingleEmail(from, subject, to, textBody, htmlBody)
	//	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	// todo should i initialize a new client every time?
	sg_client = sendgrid.NewSendClient("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")
	_, err := sg_client.Send(message)
	if err != nil {
		return err
	}

	return nil
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrcSB(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func attemptReconnectToDB() {
	// todo
}

func init() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "your-password"
		dbname   = "connect"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	accountSid := "AC9fc6898e0cad09c7b2326068a711efd8"
	authToken := "c086181d5f46faeafc6d5951ca9e58f1"
	twilio = gotwilio.NewTwilioClient(accountSid, authToken)

	sg_client = sendgrid.NewSendClient("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")

	rd_client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rd_client.Ping().Result()
	fmt.Println("redis pong? : " + pong, err)

	//_, err = db.Exec("INSERT INTO users (username, password_hash, name, email, phone_number) VALUES ('danny', 'rfewedvt', 'dann the mann', 'daf281@aol.com', '3013018023')")
	//if err != nil {
	//	panic(err)
	//}

}