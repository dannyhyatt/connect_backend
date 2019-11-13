package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
)

func isInts(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func register(c *gin.Context) {
	password := c.PostForm("password") // todo hash passwords
	name := c.PostForm("full_name")
	email := c.PostForm("email")
	phoneNumber := c.PostForm("phone_number")

	fmt.Printf(`
	registering user: 
	name %s
	email %s
	phone %s
	password %s
`, name, email, phoneNumber, password);

	if phoneNumber == "" || email == "" || password == "" || name == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success" : false,
			"error" : "Please fill out the whole form",
		})
	}
	// todo validate stuff here
	if strings.Contains(email, " ") {
		c.JSON(http.StatusNotAcceptable, gin.H {
			"success" : false,
			"error" : "Bad characters in your email address.",
		});
		return
	}

	if !isInts(phoneNumber) {
		c.JSON(http.StatusNotAcceptable, gin.H {
			"success" : false,
			"error" : "Invalid phone number.",
		});
		return
	}

	// todo fix this dumbass sql injection
	query := "INSERT INTO users (password_hash, name, email, phone_number) VALUES ('" + password + "', '" + name + "', '" + email + "', '" + phoneNumber + "')"

	// probably could do the below with Sprintf
	fmt.Printf("query: %s\n\n", query)
	_, err := db.Exec(query)
	fmt.Println("RRPRPRR")

	if err != nil {

		if strings.Contains(err.Error(), "violates unique constraint") {
			if strings.Contains(err.Error(), "users_email_key") {
				c.JSON(http.StatusNotAcceptable, gin.H {
					"success" : false,
					"error" : "The email \"" + email + "\" is already in use.",
				});
				return
			}
			if strings.Contains(err.Error(), "users_phone_number_key") {
				c.JSON(http.StatusNotAcceptable, gin.H {
					"success" : false,
					"error" : "The phone number \"" + phoneNumber + "\" is already in use.",
				});
				return
			}
		}

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

	rand.Seed(time.Now().Unix())

	phoneCode := strconv.Itoa(rand.Intn(8999) + 1000)
	emailCode := randStringBytesMaskImprSrcSB(24)

	fmt.Println("setting "+ email + PHONE_CODE_REDIS + " to " + phoneCode + " in redis")
	err = rdClient.Set(email + PHONE_CODE_REDIS, phoneCode, time.Minute * 30).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"success" : false,
			"error" : "Internal Server Error",
		})
		return
	}
	fmt.Println("^success")

	err = rdClient.Set(email + EMAIL_CODE_REDIS, emailCode, time.Hour * 30).Err() // set to "VERIFIED" when verified for both
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
	fmt.Println("sending \"" + message + "\" from " + from + " to " + phoneNumber)
	smsResponse, exception, err := twilio.SendSMS(from, to, message, "", "")
	fmt.Println("sms, exception, err:")
	fmt.Println(smsResponse)
	fmt.Println(exception)
	fmt.Println(err)
	if exception != nil {
		fmt.Println("ERROR SENDING EMAIL")
		c.JSON(http.StatusBadRequest, gin.H {
			"success" : false,
			"error" : "Error sending text. Try using another phone number?",
		})
		return
	}
	err = sendEmail(email, "Your Connect Registration Code",
		"Please go to the URL here to verify your email:\n http://"+HOSTED_URL+"api/attemptVerifyEmail/"+emailCode,
		"<img src=\"" + HOSTED_URL + "/EMAIL_LOGO.PNG\"><br><h1>Welcome to Connect!<h1> <h6>Click <a href=\"" + HOSTED_URL + "api/attemptVerifyEmail/" + email + "/" + emailCode + "\">here</a> to verify your account.</h6>",
		"Danny", "Danny from Connect", name)
	if err != nil {
		fmt.Println("ERROR SENDING EMAIL")
		c.JSON(http.StatusBadRequest, gin.H {
			"success" : false,
			"error" : "Error sending email. Try using another email address?",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"success" : true,
	})

}

func checkEmail(c *gin.Context) {

	// todo if username is taken but account is not verified check if the verification has timed out
	// and if so remove the account

	query := "SELECT * FROM users WHERE email='" + c.Param("email") + "'"
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
			"error" : "Email taken",
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"success" : true,
		})
	}
}

func attemptVerifyEmail(c *gin.Context) {
	item, err := rdClient.Get(c.Param("email") + EMAIL_CODE_REDIS).Result()
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H {
			"success" : false,
			"error" : "Internal server error. Do you have an account?",
		})
		return
	}

	//var created string
	//created, err = rd_client.Get(c.Param("email") + "_phoneCode").Result()
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
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

	if item == "VERIFIED" {
		c.JSON(http.StatusAlreadyReported, gin.H {
			"success" : false,
			"error" : "Phone number already verified.",
		})
		return
	}

	if item == c.Param("code") {
		err = rdClient.Set(c.Param("email") + EMAIL_CODE_REDIS, "VERIFIED", time.Hour * 24).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"success" : false,
				"error" : "Internal server error",
			})
			return
		}

		var success bool
		success, err = emailVerifiedCheckPhoneVerified(c.Param("email"))
		if success && err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success" : true,
			})
			return
		} else {
			fmt.Print("success: ")
			fmt.Println(success)
			fmt.Println("error with " + c.Param("email") + ", error: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success" : false,
				"error" : "Internal Service Error",
			})
			return
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success" : false,
			"error" : "Wrong validation code",
		})
		return
	}
}
func attemptVerifyPhone(c *gin.Context) {
	fmt.Println("getting " + c.Param("email") + PHONE_CODE_REDIS + " in redis")
	item, err := rdClient.Get(c.Param("email") + PHONE_CODE_REDIS).Result()
	if err != nil {
		fmt.Println("redis err::::")
		fmt.Println(err)
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

	if item == "VERIFIED" {
		c.JSON(http.StatusAlreadyReported, gin.H {
			"success" : false,
			"error" : "Phone number already verified.",
		})
		return
	}

	if item == c.Param("code") {
		err = rdClient.Set(c.Param("email") + PHONE_CODE_REDIS, "VERIFIED", time.Hour * 24).Err()
		if err != nil {
			fmt.Println("wtf2")
			c.JSON(http.StatusInternalServerError, gin.H {
				"success" : false,
				"error" : "Internal server error",
			})
			return
		}

		phoneVerifiedCheckEmailVerified(c.Param("email"))

		c.JSON(http.StatusOK, gin.H{
			"success" : true,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success" : false,
			"error" : "Wrong validation code",
		})
		return
	}
}

func isValidPhoneNumber(phoneNumber string) bool {
	// todo
	return true
}

// todo check if email is verified after phone is verified
func phoneVerifiedCheckEmailVerified(email string) (bool, error) {
	item, err := rdClient.Get(email + EMAIL_CODE_REDIS).Result()
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
		return false, err
	}

	if item == "VERIFIED" {
		verifyAccount(email)
		return true, nil
	} else {
		return false, nil
	}
}

// todo check if phone is verified after email is verified
func emailVerifiedCheckPhoneVerified(email string) (bool, error) {
	item, err := rdClient.Get(email + PHONE_CODE_REDIS).Result()
	if err != nil {
		fmt.Print("error getting redis : ")
		fmt.Println(err)
		return false, err
	}

	if item == "VERIFIED" {
		return true, verifyAccount(email)
	} else {
		return false, nil
	}
}

func verifyAccount(email string) error {
	// todo
	query := "UPDATE users SET verified=true WHERE email ='" + email + "';"
	fmt.Printf("query: %s\n\n", query)
	_, err := db.Exec(query)
	fmt.Println("YusYusY")

	if err != nil {
		return err
	}

	return nil
}
