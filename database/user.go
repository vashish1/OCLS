package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateStudentDetails(email string, input map[string]interface{}) (bool,map[string]interface{}) {
	filter := bson.D{
		{"email", email},
	}
	ok, user := Find(StudentCl, email)
    if ok{
		for key,value:=range input{
			if(user[key]!=input[key]){
				done:=UpdateData(StudentCl,filter,key,value)
				if !done{
				   return false,map[string]interface{}{}
				}
			}
		}
		ok, user = Find(StudentCl, email)
	}
	return ok,user
}
func UpdateTeacherDetails(email string, input map[string]interface{}) (bool,map[string]interface{}) {
	filter := bson.D{
		{"email", email},
	}
	ok, user := Find(TeacherCl, email)
    if ok{
		for key,value:=range input{
			if(user[key]!=input[key]){
				done:=UpdateData(TeacherCl,filter,key,value)
				if !done{
				   return false,map[string]interface{}{}
				}
			}
		}
		ok, user = Find(TeacherCl, email)
	}
	return ok,user
}

func UpdateData(c *mongo.Collection, filter interface{}, key string, value interface{}) bool{
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	update := bson.D{
		{
			"$set", bson.D{
				{key, value},
			},
		},
	}
	updateResult, err := c.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return false
	}
	return true
}
