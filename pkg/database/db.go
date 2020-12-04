package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbURL = "mongodb+srv://yashi:KNT8CQsC7120GowL@cluster0-2pscc.mongodb.net/test?retryWrites=true&w=majority"

func ConnectDb() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var retry = 3
	for retry != 0 {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(DbURL))
		if err == nil {
			fmt.Println("Connected to MongoDB")
			return client
		}
		fmt.Println("Attempt %d failed", retry)
		fmt.Println("Retrying...")
		retry--
	}
	return nil
}

func StudentDb() *mongo.Collection {
	client := ConnectDb()
	cm := client.Database("E-Vidya").Collection("Student")
	return cm
}

func TeachersDb() *mongo.Collection {
	client := ConnectDb()
	cm := client.Database("E-Vidya").Collection("Teacher")
	return cm
}

func VideoDb() *mongo.Collection {
	client := ConnectDb()
	cm := client.Database("E-Vidya").Collection("Video")
	return cm
}

func RoomDb() *mongo.Collection {
	client := ConnectDb()
	cm := client.Database("E-Vidya").Collection("Rooms")
	return cm
}
func MessageDb() *mongo.Collection {
	client := ConnectDb()
	cm := client.Database("E-Vidya").Collection("Message")
	return cm
}
func FilesDb() *mongo.Collection {
	client := ConnectDb()
	cm := client.Database("E-Vidya").Collection("Files")
	return cm
}
func ChunksDb() *mongo.Collection {
	client := ConnectDb()
	cm := client.Database("E-Vidya").Collection("FileChunks")
	return cm
}

func InsertIntoDb(c *mongo.Collection, data interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	result, err := c.InsertOne(ctx, data)
	if err != nil {
		fmt.Print("couldn't insert the document.")
		return false
	}
	fmt.Print("object id of inserted Document", result)
	return true
}