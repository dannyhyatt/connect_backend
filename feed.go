package main

import (
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
