package data

import (
	"context"
	"errors"
	"time"

	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection

func ConnectDB(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	TaskCollection = client.Database("task_manager").Collection("tasks")
	return nil
}

func CreateTask(ctx context.Context, task models.Task) (*mongo.InsertOneResult, error) {
	task.ID = primitive.NewObjectID()
	if task.DueDate.IsZero() {
		task.DueDate = time.Now()
	}
	return TaskCollection.InsertOne(ctx, task)
}

func GetAllTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := TaskCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskById(ctx context.Context, id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var task models.Task
	err = TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(ctx context.Context, id string, updatedTask models.Task) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{"$set": updatedTask}

	result, err := TaskCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no task found to update")
	}
	return nil
}
func DeleteTask(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no task found to delete")
	}
	return nil
}
