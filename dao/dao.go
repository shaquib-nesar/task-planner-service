package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/shaquib-nesar/task-planner-service.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProcessDocument(collection *mongo.Collection, ctx context.Context, message []byte) (id interface{}, err error) {
	data := models.Task{}
	err = json.Unmarshal(message, &data)
	if err != nil {
		fmt.Printf("here")
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	id = res.InsertedID

	return id, err
}

func GetAllTaskByUserId(collection *mongo.Collection, ctx context.Context, userId string) []models.Task {

	filter := bson.M{"userId": userId}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}

	var list []models.Task
	for cur.Next(ctx) {
		var task *models.Task

		e := cur.Decode(&task)
		if e != nil {
			fmt.Println(e)
		}
		list = append(list, *task)
	}
	fmt.Println(list)

	return list
}

func GetTaskById(collection *mongo.Collection, ctx context.Context, id string) (task models.Task, hasData bool) {

	filter := bson.M{"_id": id}

	hasData = false
	if err := collection.FindOne(ctx, filter).Decode(&task); err != nil {
		fmt.Println(err)
		hasData = true
	}

	return task, hasData
}

func GetSubTaskListById(collection *mongo.Collection, ctx context.Context, id string) ([]string, bool) {

	task, hasData := GetTaskById(collection, ctx, id)
	return task.SubTask, hasData
}
