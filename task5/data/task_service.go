package data

import (
	"context"
	"errors"
	"task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var collection *mongo.Collection

func ConnectDB(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	collection = client.Database("task5").Collection("todo")
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	cursor.All(ctx, &tasks)
	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	var task models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	return task, err
}

func CreateTask(task models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, task)
	return err
}

func UpdateTask(id string, updatedTask models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Ensure the ID in the struct doesn't try to overwrite the immutable _id
	updatedTask.ID = "" 
	
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedTask}
	
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}


func DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return err
}
