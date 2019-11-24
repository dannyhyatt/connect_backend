package main
// 1031 lines for backend
import (
	"database/sql"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sfreiberg/gotwilio"
	"net/http"
)

const (
	HOSTED_URL = "http://96.241.121.23:3000/"
	EMAIL_CODE_REDIS = "_emailCode"
	PHONE_CODE_REDIS = "_phoneCode"
)

var db *sql.DB
var twilio *gotwilio.Twilio
var sgClient *sendgrid.Client
var rdClient *redis.Client

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

		api.GET("/emailAvailable/:email", checkEmail)

		//api.POST("/requestNewVerificationEmail", func(c *gin.Context) {
		//	// todo
		//})

		api.POST("/login", login)

		api.POST("/register", register)

		// do we need to send credentials to verify a phone number?
		api.GET("/attemptVerifyPhone/:email/:code", attemptVerifyPhone)

		api.GET("/attemptVerifyEmail/:email/:code", attemptVerifyEmail)

		api.POST("/post/id/:id", lookupPostById)

		api.POST("/search/:query", searchByName)

		api.POST("/isFollowing", isFollowingHandler)

		api.POST("/follow", follow)

		api.POST("/unfollow", unfollow)

		api.POST("/feed", feed)

		api.POST("/profile", func(c *gin.Context) {
			validSession, err := verifySession(c.PostForm("id"), c.PostForm("session_id"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success" : false,
					"error" : "Internal server error. Try logging in again.",
				})
				return
			}
			if !validSession {
				c.JSON(http.StatusUnauthorized, gin.H{
					"success" : false,
					"error" : "Invalid session. Try logging in again.",
				})
				return
			}

			query := `SELECT email, name, phone_number FROM users WHERE id=$1;`;

			rows, err := db.Query(query, c.PostForm("id"))

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success" : false,
					"error" : "Internal server error",
				})
				return
			}

			if !rows.Next() {
				c.JSON(http.StatusNotFound, gin.H{
					"success" : false,
					"error" : "User not found",
				}) // this should never happen
				return
			}
			var email, name, phoneNumber string
			err = rows.Scan(&email, &name, &phoneNumber)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success" : false,
					"error" : "Internal server error",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"success" : true,
				"name" : name,
				"email" : email,
				"phone" : phoneNumber,
			})
			return

		})


		// for charities now
		api.POST("charity/register", charityRegisterer)

	}

//	router.NoRoute()

	// Start and run the server
	print(router.Run("0.0.0.0:3000"))
}