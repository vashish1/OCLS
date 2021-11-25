package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/vashish1/OCLS/models"
	"github.com/vashish1/OCLS/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//solve the problem of same classname
func InsertClass(input models.Class, email string) (bool, string) {
	input.Code = utility.SHA256ofstring(input.Subject)[0:6] + email[2:5]
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

func UpdataClassData(code, email, name string) bool {
	//Step 1 :add student email to class
	input := models.List{
		Name:  name,
		Email: email,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"code", code},
	}
	update := bson.M{"$push": bson.M{"studentlist": input}}
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
	update = bson.M{"$push": bson.M{"classcode": code}}
	updateResult, err = StudentCl.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return false
	}
	fmt.Println("Added %s code to student %s", code, email)
	return true
}

func InsertAssignment(desc, t, file, class, email, name string) bool {
	date, _ := time.Parse("2006-01-02T15:04", t)
	var data = models.Assignment{
		ID:          utility.GenerateUUID(),
		Classcode:   class,
		Description: desc,
		Name:        name,
		Type:        models.Type_Written,
		File: models.Written{
			FileName:    "https://storage.googleapis.com/batbuck/" + file,
			Submissions: []models.Submission{},
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

func InsertMcq(input models.Mcq, t, code, desc, email, name string) bool {
	date, _ := time.Parse("2006-01-02T15:04", t)
	input.Soln = []models.Submission{}
	var data = models.Assignment{
		ID:          utility.GenerateUUID(),
		Classcode:   code,
		Name:        name,
		Description: desc,
		Type:        models.Type_Mcq,
		Form:        input,
		Date:        date,
	}
	ok := Insert(AssignmentCl, data)
	if ok {
		if err := UpdateTeacher(email, "assignment", data.ID); err == nil {
			return true
		}
	}
	return false
}

func InsertMcqSubmission(id int, ans []string, email, name string) bool {
	date := time.Now().Format("2006-01-02 15:04:05")
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
	score := 0
	ExpAns := data.Form.Answers
	for i, x := range ExpAns {
		if ans[i] == x {
			score++
		}
	}
	input := models.Submission{
		Name:      name,
		Email:     email,
		Timestamp: date,
		Score:     score,
	}
	ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	update := bson.M{"$push": bson.M{"form.submissions": input}}
	updateResult, err := AssignmentCl.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return false
	}
	return true

}
func InsertSubmission(id, email, name, filename string) bool {

	date := time.Now().Format("2006-01-02 15:04:05")
	var sub = models.Submission{
		Email:     email,
		Name:      name,
		Timestamp: date,
		FileName:  "https://storage.googleapis.com/batbuck/" + filename,
	}
	id_value, _ := strconv.Atoi(id)
	fmt.Println(id, " ", id_value)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"id", id_value},
	}
	update := bson.M{"$push": bson.M{"file.submissions": sub}}
	updateResult, err := AssignmentCl.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return false
	}
	return true
}

func GetSubmissions(id int) (error, []models.Submission, int) {
	var data models.Assignment
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"id", id},
	}
	err := AssignmentCl.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return err, nil, -1
	}
	// fmt.Println(data)
	if data.Type == models.Type_Written {
		return nil, data.File.Submissions, data.Type
	}
	return nil, data.Form.Soln, data.Type
}

func GetStudentList(class_code string) (error, []models.List) {
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
	// fmt.Println(data)
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

func GetAllClass(email string, user_type int) (bool, []models.Class) {
	var codes []string
	if user_type == 1 {
		var data models.Teacher
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		filter := bson.D{
			{"email", email},
		}
		err := TeacherCl.FindOne(ctx, filter).Decode(&data)
		if err != nil {
			fmt.Println(err)
			return false, []models.Class{}
		}
		// fmt.Println(data)
		codes = data.Class
	} else {
		var data models.Student
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		filter := bson.D{
			{"email", email},
		}
		err := StudentCl.FindOne(ctx, filter).Decode(&data)
		if err != nil {
			fmt.Println(err)
			return false, []models.Class{}
		}
		// fmt.Println(data)
		codes = data.ClassCode
	}
	fmt.Println(codes)

	var result []models.Class
	for _, c := range codes {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		var data models.Class
		filter := bson.D{
			{"code", c},
		}
		err := ClassCl.FindOne(ctx, filter).Decode(&data)
		// fmt.Println(data)
		if err != nil {
			fmt.Println(err)
		} else {
			result = append(result, data)
		}
	}
	return true, result
}

func GetAllAnnouncement(class string) (bool, []map[string]interface{}) {
	filter := bson.D{
		primitive.E{Key: "classcode", Value: class},
	}
	err, data := FindAll(AnnouncementCl, filter)
	if err != nil {
		return false, []map[string]interface{}{}
	}
	return true, data
}

func GetAllAssignment(class string) (bool, []map[string]interface{}) {
	filter := bson.D{
		primitive.E{Key: "classcode", Value: class},
	}
	err, data := FindAll(AssignmentCl, filter)
	if err != nil {
		return false, []map[string]interface{}{}
	}
	return true, data
}

