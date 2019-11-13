package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password") // todo hash passwords

	query := "SELECT verified FROM users WHERE email='" + email + "' AND password_hash='" + password + "';"

	fmt.Println("query for db is " + query);
	var verified bool
	row := db.QueryRow(query)
	err := row.Scan(&verified)
	if err == sql.ErrNoRows {
		// todo this is firing and i don't know why
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "Incorrect password",
		})
		fmt.Println("No rows were returned!")
		return
	} else if err == nil {
		fmt.Print("email: " + email + ", verified: ")
		fmt.Println(verified)
	} else {
		fmt.Println("error logging in " + email + ": ")
		fmt.Println(err)
	}

	if !verified {
		fmt.Println("email: " + email + " is not verified")
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error": "Please verify your email",
		})
		return
	}

	//fmt.Println("email: " + email + " is verified")

	sessionId := randStringBytesMaskImprSrcSB(16)
	// err 2 incase it doesn't become null if first was error
	err2 := rdClient.Set(email + "_sessId", sessionId, time.Hour * 2).Err()
	if err2 != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"error" : "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"session_id" : sessionId,
	})
	return

}

func verifySession(email, sessionId string) (bool, error) {

	correctSessionId, err := rdClient.Get(email + "_sessId").Result()
	if err != nil {
		return false, err
	}

	return correctSessionId == sessionId, nil
}