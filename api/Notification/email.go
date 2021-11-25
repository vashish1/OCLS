package Notification

import (
	"encoding/json"
	"fmt"
	"os"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/vashish1/OCLS/database"
	"github.com/vashish1/OCLS/models"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/robfig/cron.v2"
)

func init() {
	fmt.Println("cron job working")
	c := cron.New()
	c.AddFunc("@every 48h", FilteredList)
	c.Start()
}

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
				Subject:  "New Assignment added in "+" ",
				TextPart: "Dear" +mail.Name+ "You have an assignment due, complete it before the date :" + date,
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
func FilteredList() {
	err, allclasses := database.FindAll(database.ClassCl, bson.D{{}})
	if err != nil {
		fmt.Println(err)
	}
	set1 := marshalClass(allclasses)
	for _, class := range set1 {
		list := class.StudentList
		_, assignment := database.GetAllAssignment(class.Code)
		set2 := marshalAssignment(assignment)
		process(set2, list)
	}
}

func process(set2 []models.Assignment, list []models.List) {
	for _, assign_value := range set2 {
		var target []models.Submission
		if assign_value.Type == models.Type_Mcq {
			target = assign_value.Form.Soln

		} else {
			target = assign_value.File.Submissions
		}
		for _, x := range list {
			ok := check(x.Email, target)
			if !ok {
				SendRemainder(assign_value, x)
			}
		}
	}
}

func check(email string, target []models.Submission) bool {
	for _, x := range target {
		if x.Email == email {
			return true
		}
	}
	return false
}

func marshalClass(input []map[string]interface{}) []models.Class {
	var result []models.Class
	for _, data := range input {
		res, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return []models.Class{}
		}
		var temp models.Class
		jsonRes := json.Unmarshal(res, &temp)
		if jsonRes != nil {
			fmt.Println(err)
			return []models.Class{}
		}
		result = append(result, temp)
	}
	return result
}

func marshalAssignment(input []map[string]interface{}) []models.Assignment {
	var result []models.Assignment
	for _, data := range input {
		res, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return []models.Assignment{}
		}
		var temp models.Assignment
		jsonRes := json.Unmarshal(res, &temp)
		if jsonRes != nil {
			fmt.Println(err)
			return []models.Assignment{}
		}
		result = append(result, temp)
	}
	return result
}

func SendRemainder(values models.Assignment, user models.List) {
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: from,
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: user.Email,
					Name:  user.Name,
				},
			},
			Subject:  "Assignment Due",
			TextPart: "Dear" +user.Name+ " You have an assignment due, complete it before the date :" + values.Date.Format("2006-01-02 15:04:05"),
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := client.SendMailV31(&messages)
	if err != nil {
		fmt.Println("error while sending mail", err)
		return 
	}
	//fetch all classes
    fmt.Println("email sent to",user.Name,user.Email)
	//fetch all assignments
	//seperate the students who have not submitted the assignment

}
