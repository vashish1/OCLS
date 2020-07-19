package teacher

import (
	"context"
	"fmt"
	"time"

	"github.com/vashish1/OnlineClassPortal/pkg/database"
	"github.com/vashish1/OnlineClassPortal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

var db1 = database.TeachersDb()

func Exist(email, pass string) (models.Student, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var data models.Student
	filter := bson.D{
		{"email", email},
	}
	err := db1.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return data, false
	}
	return data, true
}