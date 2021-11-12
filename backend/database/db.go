package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/vashish1/OCLS/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Colections in database
var StudentCl, TeacherCl, ClassCl *mongo.Collection

func init() {
	StudentCl, TeacherCl, ClassCl = Createdb()

}

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Collection, *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// db := os.Getenv("DbURL")
	dbURL := "mongodb+srv://yashi:Doraemon&1@cluster0.2pscc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		dbURL,
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
func UserExists(email string) (bool, map[string]interface{}) {
	if ok, user := Find(StudentCl, email); ok {
		return true, user
	} else if ok, user := Find(TeacherCl, email); ok {
		return true, user
	}
	return false, map[string]interface{}{}
}

func Insertintodb(data map[string]interface{}) (bool, error) {
	var ok bool
	//If user type is student insert in student database else in teacher database
	if (int)(data["type"].(float64)) == models.Type_Student {
		ok = Insert(StudentCl, data)
	} else {
		ok = Insert(TeacherCl, data)
	}
	if ok {
		return true, nil
	}

	return false, errors.New("Error while inserting into database")
}

func UpdateTeacher(email, code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"email", email},
	}
	update := bson.M{"$push": bson.M{"class": code}}
	updateResult, err := TeacherCl.UpdateOne(ctx, filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		fmt.Println(err)
		return errors.New("error while inserting updating the teacher data")
	}
	return nil
}