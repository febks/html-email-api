package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func SendEmail(to []string, cc []string, subject, htmlBody string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	authEmail := os.Getenv("CONFIG_AUTH_EMAIL")
	authPassword := os.Getenv("CONFIG_AUTH_PASSWORD")
	senderName := os.Getenv("CONFIG_SENDER_NAME") + " <" + authEmail + ">"

	// create Headers
	headers := make(map[string]string)
	headers["From"] = senderName
	headers["To"] = strings.Join(to, ",")
	if len(cc) > 0 {
		headers["Cc"] = strings.Join(cc, ",")
	}
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	// create email body
	var msg strings.Builder
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n" + htmlBody)

	auth := smtp.PlainAuth("", authEmail, authPassword, smtpHost)
	smtpAddr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	allRecipients := append(to, cc...)

	err := smtp.SendMail(smtpAddr, auth, authEmail, allRecipients, []byte(msg.String()))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}
