package main

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sfreiberg/gotwilio"
)

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

	sgClient = sendgrid.NewSendClient("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")

	rdClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdClient.Ping().Result()
	fmt.Println("redis pong? : " + pong, err)

	//_, err = db.Exec("INSERT INTO users (username, password_hash, name, email, phone_number) VALUES ('danny', 'rfewedvt', 'dann the mann', 'daf281@aol.com', '3013018023')")
	//if err != nil {
	//	panic(err)
	//}

}
