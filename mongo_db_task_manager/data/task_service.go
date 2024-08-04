// data/task_service.go
package data

import (
	"context"
	"errors"
	"time"

	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return tasks, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var task models.Task
		if err = cursor.Decode(&task); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	if err = cursor.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetTaskByID(id string) (*models.Task, error) {
	var task models.Task
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func CreateTask(task models.Task) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	task.ID = result.InsertedID.(primitive.ObjectID)
	return &task, nil
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedTask}
	// opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	da ,err := TaskCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	if da.MatchedCount == 0 {
		return nil, errors.New("task not found")
	}
	

	return &updatedTask, nil
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return errors.New("task not found")
	}

	return nil
}
func DeleteAllTasks() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := TaskCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return errors.New("task not found")
	}

	return nil
}