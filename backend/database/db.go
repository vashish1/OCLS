package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Colections in database
var studentCl, teacherCl, classCl *mongo.Collection

func init() {
	studentCl, teacherCl, classCl = Createdb()

}

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Collection, *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := os.Getenv("DbURL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		db,
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	studentdb := client.Database("OCLS").Collection("Student")
	teacherdb := client.Database("OCLS").Collection("Teacher")
	classdb := client.Database("OCLS").Collection("Class")

	return studentdb, teacherdb, classdb
}

func Insertintodb(data interface{})(bool,error){
  return true,nil;
}