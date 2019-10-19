package main

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sfreiberg/gotwilio"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
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

		api.POST("/requestNewVerificationEmail", func(c *gin.Context) {
			// todo
		})

		api.POST("/login", login)

		api.POST("/register", register)

		// do we need to send credentials to verify a phone number?
		api.GET("/attemptVerifyPhone/:email/:code", attemptVerifyPhone)

		api.GET("/attemptVerifyEmail/:email/:code", attemptVerifyEmail)

		api.POST("/post/id/:id", lookupPostById)

		api.POST("/search/:query", func(c *gin.Context) {
			if c.Param("query") == nil || c.Param("query") == "" {
				c.JSON(http.StatusNoContent, gin.H{
					"success" : false,
					"error" : "No search query provided"
				})
				return
			}

			query := "SELECT * FROM charities WHERE LOWER(short_name) LIKE '%' || LOWER('$1') || '%' OR LOWER(long_name) LIKE '%' || LOWER('$1') || '%' LIMIT 5;"
		})

		// todo change this to return list of post id's to fetch
		api.POST("/feed", func(c *gin.Context) {
			validSession, err := verifySession(c.PostForm("email"), c.PostForm("sessId"))
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

			// TODO WHERE I REALLY LEFT OFF

			// todo get the charities the user follows

			// todo get the most recent posts from the ones he follows

			// todo id returns 0 for some reason
			id, err := getUserIdFromEmail(c.PostForm("email"))
			fmt.Print("ID for email " + c.PostForm("email") + " is ")
			fmt.Println(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success" : false,
					"error" : "Internal server error. Anna oop.",
				})
				return
			}

			query := `select title, content, thumbnail, post_time, last_edit from (
							select distinct on (post_time) *
							from charity_posts
							order by post_time
						) t WHERE charity_id in (SELECT charity_id FROM  followers WHERE user_id=$1)
						order by post_time limit 3;`;

			rows, err := db.Query(query, id)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success" : false,
					"error" : "Database error. Anna oop 2.",
				})
				return
			}

			var items [3]gin.H
			var title, content, thumbnail string
			var postDate, lastEdit time.Time

			for i := 0; i < 3; i++ {
				if !rows.Next() { break }
				rows.Scan(&title, &content, &thumbnail, &postDate, &lastEdit)
				items[i] = gin.H{
					"title" : title,
					"content" : content,
					"thumbnail" : thumbnail,
					"postDate" : postDate.Unix(),
					"lastEdit" : postDate.Unix(),
				}
			}

			c.JSON(http.StatusOK, gin.H{
				"success" : true,
				"items" : items,
			})
		})


		// for charities now
		api.POST("charity/register", charityRegisterer)

	}

//	router.NoRoute()

	// Start and run the server
	print(router.Run("0.0.0.0:3000"))
}