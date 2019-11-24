package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		rows.Close()
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
	rows.Close()
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

func follow(c *gin.Context) {

	fmt.Println(c.Accepted)
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

	validSession, err := verifySession(id, sessionId)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
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

	if charityId == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"success" : false,
			"error" : "Charity not found.",
		})
		return
	}

	//query := "INSERT INTO followers (user_id, charity_id) VALUES ($1, $2);"
	query := "INSERT INTO followers (user_id, charity_id) VALUES (" + id + ", " + charityId + ");"
	fmt.Println("query for follow: " + query)

	_, err = db.Exec(query)

	if err != nil {
		fmt.Println("error: ")
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"error" : "Unknown error", // todo
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
	})
	return
}

func unfollow(c *gin.Context) {

	id := c.PostForm("id")
	sessionId := c.PostForm("session_id")
	charityId := c.PostForm("charity_id")

	validSession, err := verifySession(id, sessionId)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
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

	if charityId == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"success" : false,
			"error" : "Charity not found.",
		})
		return
	}

	//query := "DELETE FROM followers WHERE user_id=$1 AND charity_id=$2;"
	query := "DELETE FROM followers WHERE user_id=" + id + " AND charity_id=" + charityId + ";"

	fmt.Println("query is: " + query)
	_, err = db.Exec(query)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"error" : "Unknown error", // todo
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
	})
	return
}

func getLastFiveFromViewN(id string, n int) ([]string, error) {

	// todo renae to charity id
	query := `SELECT charity_post_id FROM views WHERE user_id=$1 ORDER BY viewed_at DESC LIMIT 5 OFFSET $2;`;

	rows, err := db.Query(query, id, n)

	if err != nil {
		return nil, err
	}

	var items = make([]string, 5)
	cutoff := 5

	for i := 0; i < 5; i++ {
		if !rows.Next() { break }
		rows.Scan(&items[i]) // todo handle error
		fmt.Println("just scanned ", items[i])
		if items[i] == "" {
			cutoff = i - 1 // handle if i is 0 index out of bounds
		}
	}

	return items[:cutoff], nil;
}

func feed(c *gin.Context) {
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

	offset := 0
	if c.PostForm("offset") != "" {
		if c.PostForm("offset") == "new" {
			query := `INSERT INTO views (user_id, charity_post_id, viewed_at) VALUES ($1, unnest(array(
					select id from (
						select distinct on (last_edit) *
							from charity_posts
							order by last_edit desc
							) t WHERE charity_id in (SELECT charity_id FROM  followers WHERE user_id=$1) AND id NOT IN (SELECT charity_post_id FROM views WHERE user_id=$1) -- that's untested
							order by last_edit limit 5 offset $2)
						), now());`
			fmt.Println("query for follow: " + query)

			_, err = db.Exec(query, c.PostForm("id"), offset)

			if err != nil {
				fmt.Println("1error: ")
				fmt.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"success" : false,
					"error" : "Unknown error", // todo
				})
				return
			}
		} else {
			offset, err = strconv.Atoi(c.PostForm("offset"))
			if err != nil {
				offset = 0
			}
		}
	}

	postIds, err := getLastFiveFromViewN(c.PostForm("id"), offset)
	if err != nil {
		fmt.Printf("error geetting last 5 \n\n")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"error" : "Internal server error. Try logging in again.",
		})
		return
	}

	fmt.Println(postIds)

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"feed" : postIds,
	})
	return

	//// todo id returns 0 for some reason
	//id, err := getUserIdFromEmail(c.PostForm("email"))
	//fmt.Print("ID for email " + c.PostForm("email") + " is ")
	//fmt.Println(id)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"success" : false,
	//		"error" : "Internal server error. Anna oop.",
	//	})
	//	return
	//}
	//
	//query := `INSERT INTO views (user_id, charity_post_id, viewed_at) VALUES ($1, unnest(array(
	//			select id from (
	//					 select distinct on (last_edit) *
	//					 from charity_posts
	//					 order by last_edit desc
	//				 ) t WHERE charity_id in (SELECT charity_id FROM  followers WHERE user_id=$1) AND id NOT IN (SELECT charity_post_id FROM views WHERE user_id=$1) -- that's untested
	//				 order by last_edit limit 2)
	//			), now());`;
	//
	//rows, err := db.Exec(query, id)
	//
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"success" : false,
	//		"error" : "Database error. Anna oop 2.",
	//	})
	//	return
	//}
	//
	//var items [5]gin.H
	//var title, content, thumbnail string
	//var postDate, lastEdit time.Time
	//
	//for i := 0; i < 3; i++ {
	//	if !rows.Next() { break }
	//	rows.Scan(&title, &content, &thumbnail, &postDate, &lastEdit)
	//	items[i] = gin.H{
	//		"title" : title,
	//		"content" : content,
	//		"thumbnail" : thumbnail,
	//		"postDate" : postDate.Unix(),
	//		"lastEdit" : postDate.Unix(),
	//	}
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"success" : true,
	//	"items" : items,
	//})
}