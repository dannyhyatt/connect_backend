package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password") // todo hash passwords

	query := "SELECT id, verified FROM users WHERE email='" + email + "' AND password_hash='" + password + "';"

	fmt.Println("query for db is " + query);
	fmt.Println("password: " + password)
	var verified bool
	var id_int int64
	row := db.QueryRow(query)
	err := row.Scan(&id_int, &verified)
	id := strconv.FormatInt(id_int, 10)
	fmt.Println("id is " + id)
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
	err2 := rdClient.Set(id + "_sessId", sessionId, time.Hour * 2).Err()
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
		"id" : id,
	})
	return

}

func verifySession(id, sessionId string) (bool, error) {

	correctSessionId, err := rdClient.Get(id + "_sessId").Result()
	fmt.Printf("verifying %s with key %s, error:\n", id + "_sessId", sessionId)
	fmt.Println(err)
	if err != nil {
		return false, err
	}

	return correctSessionId == sessionId, nil
}