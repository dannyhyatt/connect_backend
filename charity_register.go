package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func charityRegisterer(c *gin.Context) {
	short_name := c.PostForm("short_name")
	long_name := c.PostForm("long_name") // todo hash passwords
	description := c.PostForm("description")
	ceo := c.PostForm("ceo")
	password := c.PostForm("password")

	query := "INSERT INTO charities (short_name, long_name, description, ceo, password_hash) VALUES ('" + short_name + "', '" + long_name + "', '" + description + "', '" + ceo + "', '" + password + "')"

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
	}

	c.JSON(http.StatusOK, gin.H {
		"success" : true,
	})
}