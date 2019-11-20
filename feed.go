package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func lookupPostById(c *gin.Context) {

	query := "SELECT title, content, author_id, charity_id, thumbnail, last_edit FROM charity_posts WHERE id=$1;"
	rows, err := db.Query(query, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"error" : "Database error. Anna oop 3.",
		})
		return
	}


	var title, content, author_id, charity_id, thumbnail string
	var lastEdit time.Time
	if !rows.Next()  {
		c.JSON(http.StatusNotFound, gin.H {
			"success" : false,
			"error" : "Post not found.",
		})
		return
	}

	rows.Scan(&title, &content, &author_id, &charity_id, &thumbnail, &lastEdit)

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"title" : title,
		"content" : content,
		"author_id" : author_id,
		"charity_id" : charity_id,
		"thumbnail" : thumbnail,
		"last_edit" : lastEdit.Unix(),
	})
	return

}

func isFollowingHandler(c *gin.Context) {

	id, a := c.GetPostForm("id")
	sessionId, b := c.GetPostForm("session_id")
	charityId, d := c.GetPostForm("charity_id")
	if !a || !b || !d {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success" : false,
			"error" : "Unspecified parameters",
		})
		return
	}

	fmt.Printf("just received id: %s, session id: %s, charity id: %s\n",id, sessionId, charityId)

	validSession, err := verifySession(id, sessionId)
	if err != nil {
		fmt.Println("err for validating session")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"error" : "Internal server error. Try logging in again.",
		})
		return
	}
	if !validSession {
		//fmt.Println("invalid session when requesting charity: " + c.PostForm("charity_id"))
		c.JSON(http.StatusUnauthorized, gin.H{
			"success" : false,
			"error" : "Invalid session. Try logging in again.",
		})
		return
	}
	if c.PostForm("charity_id") == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"success" : false,
			"error" : "Charity not found.",
		})
		return
	}

	following, err := isFollowing(id, charityId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"error" : "Database error",
		})
		fmt.Print("error: ")
		fmt.Println(err)
		return
	}

	fmt.Printf("\nfollowing: %v, user: %s, charity: %s\n", following, id, charityId)

	c.JSON(http.StatusOK, gin.H {
		"success" : true,
		"following" : following,
		"charity_id" : charityId,
	})
	return
}

func isFollowing(user_id string, charity_id string) (bool, error) {
	query := "SELECT * FROM followers WHERE user_id=$1 AND charity_id=$2;"
	rows, err := db.Query(query, user_id, charity_id)

	if err != nil {
		return false, err
	}

	if rows.Next() {
		if rows.Close() != nil {
			fmt.Println("error closing db")
			return false, err
		}
		fmt.Println("successfully closed db")
		return true, nil
	} else {
		if rows.Close() != nil {
			fmt.Println("error closing db")
			return false, err
		}
		fmt.Println("successfully closed db")
		return false, nil
	}

}

func searchByName(c *gin.Context) {
	if c.Param("query") == "" {
		c.JSON(http.StatusNoContent, gin.H{
			"success" : false,
			"error" : "No search query provided",
		})
		return
	}

	const itemLimit = 5
	var searchStr string
	searchStr = c.Param("query")

	query := "SELECT id, short_name, long_name, description, profile_url FROM charities WHERE LOWER(short_name) LIKE '%' || LOWER($1) || '%' OR LOWER(long_name) LIKE '%' || LOWER($1) || '%' LIMIT $2;"
	rows, err := db.Query(query, searchStr, itemLimit)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"error" : "Database error. Anna oop 2.",
		})
		return
	}

	var id, shortName, longName, description, profile_url string
	var results [itemLimit]gin.H

	for i := 0; rows.Next(); i++ {
		rows.Scan(&id, &shortName, &longName, &description, &profile_url)
		results[i] = gin.H{
			"id" : id,
			"shortName" : shortName,
			"longName" : longName,
			"description" : description,
			"profileUrl" : profile_url,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"results" : results,
	})
	return
}