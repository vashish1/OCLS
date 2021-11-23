package Notification

import (
	"fmt"
	"os"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/vashish1/OCLS/database"
)

var key, pass = os.Getenv("SMTP_KEY"), os.Getenv("SMTP_PASS")

// var key = "3c00f30e08727673f60cbab04ca3e4a2"
// var pass = "3658491e742f5fa9d2f54fa1d76a556f"

var from = &mailjet.RecipientV31{
	Email: "vashishtiv@gmail.com",
	Name:  "Yashi Gupta",
}

var client = mailjet.NewMailjetClient(key, pass)

//SendWelcomeEmail func
func SendEmail(class_code, date string) bool {
	err, list := database.GetStudentList(class_code)
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, mail := range list {
		messagesInfo := []mailjet.InfoMessagesV31{
			mailjet.InfoMessagesV31{
				From: from,
				To: &mailjet.RecipientsV31{
					mailjet.RecipientV31{
						Email: mail.Email,
						Name:  mail.Name,
					},
				},
				Subject:  "Assignment Due",
				TextPart: "Dear student\n" + "You have an assignment due, complete it before the date :" + date,
			},
		}
		messages := mailjet.MessagesV31{Info: messagesInfo}
		_, err := client.SendMailV31(&messages)
		if err != nil {
			fmt.Println("error while sending mail", err)
			return false
		}

	}
	return true
}
