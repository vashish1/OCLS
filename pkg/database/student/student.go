package student

import (
	"context"
	"fmt"
	"time"

	"github.com/vashish1/OnlineClassPortal/pkg/database"
	"github.com/vashish1/OnlineClassPortal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

var db = database.StudentDb()

func Exist(email, pass string) (models.Student, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var data models.Student
	filter := bson.D{
		{"email", email},
		{"freeze", false},
	}
	err := db.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return data, false
	}
	return data, true
}

func IsAvailable(uid string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"uid", uid},
		{"freeze", false},
	}
	err := db.FindOne(ctx, filter)
	if err.Err() == nil {
		fmt.Println("Same Uid exists", err)
		return true
	}
	return false
}
