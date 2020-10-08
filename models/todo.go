package models

import (
	"context"
	"fmt"
	config "todolist2/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

var (
	TodoList map[string]*Todo
)

type Todo struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title"`
	Note  string             `bson:"note"`
	Date  string             `bson:"date"`
}

func AddTodo(t Todo) string {

	res, err := config.Database.Collection("Todo").InsertOne(context.Background(), t)

	if err != nil {
		fmt.Println("error")
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)

	fmt.Println(oid)

	return oid.Hex()
}

func GetAllTodos() []Todo {

	cursor, _ := config.Database.Collection("Todo").Find(context.TODO(), bson.M{})
	var todos []Todo

	for cursor.Next(context.TODO()) {
		data := Todo{}
		cursor.Decode(&data)
		todos = append(todos, data)
	}

	return todos
}

func DeleteTodo(uid string) {
	fmt.Println(uid)
	ID, _ := primitive.ObjectIDFromHex(uid)
	config.Database.Collection("Todo").DeleteOne(context.Background(), bson.M{"_id": ID})
}

func EditTodo(t Todo) {
	config.Database.Collection("Todo").UpdateOne(
		context.Background(),
		bson.M{"_id": t.ID},
		bson.M{"$set": bson.M{
			"title": t.Title,
			"note":  t.Note,
			"date":  t.Date,
		}},
	)
}
