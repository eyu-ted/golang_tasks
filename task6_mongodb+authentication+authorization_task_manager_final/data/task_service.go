// data/task_service.go
package data

import (
	"context"
	"errors"
	"fmt"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task_manager/middleware"


	
)

func IsAdmin(userID primitive.ObjectID) bool {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := UserCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return false
	}

	return user.Role == "admin"
}

func GetAllTasks(userID primitive.ObjectID) ([]models.Task, error) {
    var tasks []models.Task
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	filter := bson.M{}
	if !IsAdmin(userID) {
		filter = bson.M{"ownerid": userID}
	}
	fmt.Println(filter)

    cursor, err := TaskCollection.Find(ctx, filter)
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

func GetTaskByID(id string, userID primitive.ObjectID) (*models.Task, error) {
    var task models.Task
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, errors.New("invalid ID")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	filter := bson.M{"_id": objID}
	if !IsAdmin(userID) {
    	filter = bson.M{"_id": objID, "ownerid": userID}
}
    err = TaskCollection.FindOne(ctx, filter).Decode(&task)
    if err != nil {
        return nil, errors.New("task not found")
    }

    return &task, nil
}



func CreateTask(user *middleware.Claims, task models.Task) (*models.Task, error) {
    fmt.Println(user.UserID)

    task.OwnerID = user.UserID

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    result, err := TaskCollection.InsertOne(ctx, task)
    if err != nil {
        return nil, err
    }

    task.ID = result.InsertedID.(primitive.ObjectID)
    return &task, nil
}
func UpdateTaskUsers(user *middleware.Claims, task models.Task) (*models.Task, error) {
    // Print the user ID for debugging purposes
    fmt.Println("User ID:", user.UserID)

    // Create a context with a timeout to prevent long-running operations
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	task.OwnerID = user.UserID
	filter := bson.M{"_id": task.ID, "ownerid": user.UserID}
	update := bson.M{"$set": task}
	// opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	da ,err := TaskCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	if da.MatchedCount == 0 && !IsAdmin(user.UserID) {
		return nil, fmt.Errorf("unable to update task: either the task does not exist or you are not authorized to update it")
	}
	filter = bson.M{"_id": task.ID}
	da ,err = TaskCollection.UpdateOne(ctx, filter, update)

    return &task, nil
}
func DeleteTask(id string, userID primitive.ObjectID) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New("invalid ID")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Find the task first to verify ownership
    var task models.Task
    err = TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
    if err != nil {
        return errors.New("task not found")
    }

    // Check if the user is the owner of the task
    if task.OwnerID != userID && !IsAdmin(userID) {

        return errors.New("not authorized to delete this task")
    }

    // Proceed to delete the task
    filter := bson.M{"_id": objID}
    result, err := TaskCollection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }

    if result.DeletedCount == 0 {
        return errors.New("task not found")
    }

    return nil
}

