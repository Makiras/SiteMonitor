package utils

import (
	"log"
	"net/smtp"
)

// MailAuthInfo describes the info for smtp authorization
type MailAuthInfo struct {
	EmailAddr  string
	Pass       string // prefer AES Encrypted
	ServerAddr string
}

// SendMail Provides sending mail function
func SendMail(authinfo MailAuthInfo, to []string, msg []byte) error {
	auth := smtp.PlainAuth(
		"",
		authinfo.EmailAddr,
		authinfo.Pass, // decrypt required
		authinfo.ServerAddr,
	)

	err := smtp.SendMail(
		authinfo.ServerAddr,
		auth,
		authinfo.EmailAddr,
		to,
		msg,
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
