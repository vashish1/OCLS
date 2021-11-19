package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/vashish1/OCLS/backend/models"
	"github.com/vashish1/OCLS/backend/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//solve the problem of same classname
func InsertClass(input models.Class, email string) (bool, string) {
	input.Code = utility.SHA256ofstring(input.TeacherEmail)[0:6]
	ok := Insert(ClassCl, input)
	if ok {
		err := UpdateTeacher(email, "class", input.Code)
		if err != nil {
			fmt.Println(err)
			return false, ""
		}
		return true, input.Code
	}
	return false, ""
}

func UpdataClassData(code, email string) bool {
	//Step 1 :add student email to class
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"code", code},
	}
	update := bson.M{"$push": bson.M{"student_list": email}}
	updateResult, err := ClassCl.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return false
	}
	fmt.Println("Added %s student to class %s", email, code)

	//Step 2 :add class to student's details
	filter = bson.D{
		{"email", email},
	}
	update = bson.M{"$push": bson.M{"class_code": code}}
	updateResult, err = StudentCl.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return false
	}
	fmt.Println("Added %s code to student %s", code, email)
	return true
}

func InsertAssignment(desc, t, file, class, email string) bool {
	date, _ := time.Parse("2006-01-02T15:04", t)
	var data = models.Assignment{
		ID:          utility.GenerateUUID(),
		Class_code:  class,
		Description: desc,
		Type:        models.Type_Written,
		File: models.Written{
			FileName: "https://storage.googleapis.com/batbuck/" + file,
		},
		Date: date,
	}
	ok := Insert(AssignmentCl, data)
	if ok {
		if err := UpdateTeacher(email, "assignment", data.ID); err == nil {
			return true
		}
	}
	return false
}

func InsertMcq(input models.Mcq,t,code,desc, email string) bool{
	date, _ := time.Parse("2006-01-02T15:04", t)
	var data = models.Assignment{
		ID:          utility.GenerateUUID(),
		Class_code:  code,
		Description: desc,
		Type:        models.Type_Mcq,
		Form: input,
		Date: date,
	}
	ok := Insert(AssignmentCl, data)
	if ok {
		if err := UpdateTeacher(email, "assignment", data.ID); err == nil {
			return true
		}
	}
	return false
}
func InsertMcqSubmission(id int,ans []string,email string)bool{
	date := time.Now().Format("2006-01-02 T 15:04")
	var data models.Assignment
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"id", id},
	}
	err := AssignmentCl.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return false
	}
	score:=0
    ExpAns:=data.Form.Answers
	for i,x:=range ExpAns{
      if ans[i]==x {
		  score++
	  }
	}
   _,user:=Find(StudentCl,email)
   input := models.Submission{
		Name: user["name"].(string),
		Email: email,
		Timestamp: date,
		Score: score,
	}
	ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	update := bson.M{"$push": bson.M{"form.submissions": input}}
	updateResult, err := AnnouncementCl.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return false
	}
	return true

}
func InsertSubmission(id, email, filename string) bool {

	date := time.Now().Format("2006-01-02 T 15:04")
	var sub = models.Submission{
		Email:     email,
		Timestamp: date,
		FileName:  "https://storage.googleapis.com/batbuck/" + filename,
	}
	id_value, _ := strconv.Atoi(id)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"id", id_value},
	}
	update := bson.M{"$push": bson.M{"file.submissions": sub}}
	updateResult, err := AnnouncementCl.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return false
	}
	return true
}

func GetSubmissions(id int) (error, []models.Submission) {
	var data models.Assignment
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"id", id},
	}
	err := AssignmentCl.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return err, nil
	}
	fmt.Println(data)
	return nil, data.File.Submissions
}

func GetStudentList(class_code string) (error, []string) {
	var data models.Class
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"code", class_code},
	}
	err := ClassCl.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return err, nil
	}
	fmt.Println(data)
	return nil, data.StudentList
}

func InsertAnnouncement(input models.Announcement, email string) bool {
	input.ID = utility.GenerateUUID()
	ok := Insert(AnnouncementCl, input)
	if ok {
		err := UpdateTeacher(email, "post", input.ID)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return false
}

func GetAllClass() (bool, []map[string]interface{}) {
	filter := bson.D{{}}
	err, data := FindAll(ClassCl, filter)
	if err != nil {
		return false, []map[string]interface{}{}
	}
	return true, data
}

func GetAllAnnouncement(class string) (bool, []map[string]interface{}) {
	filter := bson.D{
		primitive.E{Key: "class_code", Value: class},
	}
	err, data := FindAll(AnnouncementCl, filter)
	if err != nil {
		return false, []map[string]interface{}{}
	}
	return true, data
}

func GetAllAssignment(class string) (bool, []map[string]interface{}) {
	filter := bson.D{
		primitive.E{Key: "class_code", Value: class},
	}
	err, data := FindAll(AssignmentCl, filter)
	if err != nil {
		return false, []map[string]interface{}{}
	}
	return true, data
}
