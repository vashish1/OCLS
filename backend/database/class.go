package database

import (
	"context"
	"fmt"
	"time"

	"github.com/vashish1/OCLS/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//solve the problem of same classname
func InsertClass(input models.Class) (bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	result, err := ClassCl.InsertOne(ctx, input)
	if err != nil {
		fmt.Print("couldn't insert the document.")
		return false, ""
	}
	Id :=((result.InsertedID).(primitive.ObjectID)).String()
     code := Id[10:16]
	return true,code
}

