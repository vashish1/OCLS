package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(c *mongo.Collection, data interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	_, err := c.InsertOne(ctx, data)
	if err != nil {
		fmt.Print("couldn't insert the document.")
		return false
	}
	return true
}

func Find(c *mongo.Collection, email string) (bool, map[string]interface{}) {
	var data map[string]interface{}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"email", email},
	}
	err := c.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return false, map[string]interface{}{}
	}
	fmt.Println(data)
	return true, data
}

// func GetData(c *mongo.Collection,key string,value interface{}) (error,map[string]interface{}){
// 	var data map[string]interface{}
// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	defer cancel()
// 	filter := bson.D{
// 		{key, value},
// 	}
// 	err := c.FindOne(ctx, filter).Decode(&data)
// 	if err != nil {
// 		return err,nil
// 	}
// 	fmt.Println(data)
// 	return nil,data
// }

