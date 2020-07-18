package student

import (
	"context"
	"time"

	"github.com/vashish1/OnlineClassPortal/pkg/database"
	"github.com/vashish1/OnlineClassPortal/pkg/helpers"
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
	}
	err := db.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return data,false
	}
	ok := helpers.ValidatePass(data.PassHash, pass)
	return data,ok
}
