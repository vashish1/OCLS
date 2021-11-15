package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/vashish1/OCLS/backend/models"
	"github.com/vashish1/OCLS/backend/utility"
	"go.mongodb.org/mongo-driver/bson"
)

//solve the problem of same classname
func InsertClass(input models.Class) (bool, string) {
	input.Code = utility.SHA256ofstring(input.TeacherEmail)[0:6]
	ok := Insert(ClassCl, input)
	if ok {
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

func InsertAssignment(desc, t, file string) (bool, int) {
	date, _ := time.Parse("2006-01-02T15:04", t)
	var data = models.Assignment{
		ID:          utility.GenerateUUID(),
		Description: desc,
		FileName:    "https://storage.googleapis.com/batbuck/"+file,
		Date:        date,
	}
	return Insert(AssignmentCl, data), data.ID
}

func InsertSubmission(id, email, filename string) bool {
	  
	date, _ := time.Parse("2006-01-02T15:04", time.Now().Format("2006-01-02T15:04") )
	var sub = models.Submission{
		Email:     email,
		Timestamp: date,
		FileName:  "https://storage.googleapis.com/batbuck/"+filename,
	}
	id_value,_ := strconv.Atoi(id)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"id",id_value},
	}
	update := bson.M{"$push": bson.M{"submissions": sub}}
	updateResult, err := AnnouncementCl.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return false
	}
	return true
}

func GetSubmissions(id int) (error,[]models.Submission){
	var data models.Assignment
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		filter := bson.D{
			{"id", id},
		}
		err := AssignmentCl.FindOne(ctx, filter).Decode(&data)
		if err != nil {
			return err,nil
		}
		fmt.Println(data)
		return nil,data.Submissions
}

func GetStudentList(class_code string) (error,[]string){
	var data models.Class
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		filter := bson.D{
			{"code", class_code},
		}
		err := ClassCl.FindOne(ctx, filter).Decode(&data)
		if err != nil {
			return err,nil
		}
		fmt.Println(data)
		return nil,data.StudentList
}