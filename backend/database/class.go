package database

import (
	"context"
	"fmt"
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

func InsertAssignment(desc,file string)(bool,int){
	var data = models.Assignment{
		ID: utility.GenerateUUID(),
		Description: desc,
		FileName:file,
	}
	return Insert(AssignmentCl,data),data.ID
}