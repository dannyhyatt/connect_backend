package main

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"math/rand"
	"strings"
	"time"
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrcSB(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func sendEmail(address, subject, textBody, htmlBody, fromPrefix, fromName, toName string) error {
	if fromPrefix == "" {
		fromPrefix = "connect"
	}
	if fromName == "" {
		fromName = "Connect"
	}
	from := mail.NewEmail(fromName, fromPrefix + "@connect.charity")
	to := mail.NewEmail(toName, address)
	if htmlBody == "" {
		htmlBody = textBody
	}
	message := mail.NewSingleEmail(from, subject, to, textBody, htmlBody)
	//	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	// todo should i initialize a new client every time?
	sgClient = sendgrid.NewSendClient("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")
	_, err := sgClient.Send(message)
	if err != nil {
		return err
	}

	return nil
}

func getUserIdFromEmail(email string) (int64, error) {
	query := "SELECT id FROM users WHERE email=$1;"
	rows, err := db.Query(query, email)

	if err != nil {
		return -1, err
	}

	//fmt.Println("rows?")

	if rows.Next() {
		//fmt.Println("success please")
	}

	//fmt.Println("what is di? ")

	var id int64
	err = rows.Scan(&id)

	if err != nil {
		return -1, err
	}

	//fmt.Sprintf("ID FOR " + email + " IS %d", id)

	return id, nil
}