package models

import (
	"context"
	"fmt"
	config "todolist2/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

var (
	UserList map[string]*User
)

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

func AddUser(t User) string {
	res, err := config.Database.Collection("User").InsertOne(context.Background(), t)

	if err != nil {
		fmt.Println("error")
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)

	fmt.Println(oid)

	return oid.Hex()
}

func GetAllUsers() []User {
	cursor, _ := config.Database.Collection("User").Find(context.TODO(), bson.M{})
	var Users []User

	for cursor.Next(context.TODO()) {
		data := User{}
		cursor.Decode(&data)
		Users = append(Users, data)
	}

	return Users
}

func DeleteUser(uid string) {
	fmt.Println(uid)
	ID, _ := primitive.ObjectIDFromHex(uid)
	config.Database.Collection("User").DeleteOne(context.Background(), bson.M{"_id": ID})
}

func EditUser(t User) {
	config.Database.Collection("User").UpdateOne(
		context.Background(),
		bson.M{"_id": t.ID},
		bson.M{"$set": bson.M{
			"name":  t.Name,
			"email": t.Email,
		}},
	)
}
