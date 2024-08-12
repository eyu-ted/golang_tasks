package repositories

import (
	"context"
	// "path/filepath"
	"tskmgr/domain"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskDataManipulator struct {
	collection *mongo.Collection
}

// GetAllTasks implements domain.TaskRepository.
func (m *TaskDataManipulator) GetAllTasks(userRole string, userID primitive.ObjectID) ([]*domain.Task, error) {
	var tasks []*domain.Task

	filter := bson.M{"userid": userID}
	if userRole == "admin" {
		filter = bson.M{}
	}

	cursor, err := m.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task domain.Task
		if err = cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

}

func NewTaskDataManipulator(collection *mongo.Collection) *TaskDataManipulator {
	return &TaskDataManipulator{
		collection: collection,
	}
}

func (m *TaskDataManipulator) StoreTask(task *domain.Task) (*domain.Task, error) {
	_, err := m.collection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (m *TaskDataManipulator) GetByTitle(title string) (*domain.Task, error) {
	var task domain.Task
	err := m.collection.FindOne(context.TODO(), bson.M{"title": title}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// func (m *TaskDataManipulator) GetAllTasks() ([]*domain.Task, error) {
// 	var tasks []*domain.Task

// 	cursor, err := m.collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.TODO())

// 	for cursor.Next(context.TODO()) {
// 		var task domain.Task
// 		if err = cursor.Decode(&task); err != nil {
// 			return nil, err
// 		}
// 		tasks = append(tasks, &task)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return tasks, nil
// }

func (m *TaskDataManipulator) GetUserTasks(userid primitive.ObjectID) ([]*domain.Task, error) {
	var tasks []*domain.Task

	cursor, err := m.collection.Find(context.TODO(), bson.M{"userid": userid})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task domain.Task
		if err = cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *TaskDataManipulator) UpdateTask(userRole string, userID primitive.ObjectID, title string, task *domain.Task) (*domain.Task, error) {

	storedTask, err := m.GetByTitle(title)
	if err != nil {
		return nil, err
	}

	if storedTask.UserId != userID && userRole != "admin" {
		return nil, errors.New("unauthorized to update this task")
	}

	// task.UserId = storedTask.UserId
	filter := bson.M{"title": task.Title}
	update := bson.M{"$set": task}

	_, err = m.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (m *TaskDataManipulator) DeleteTask(userRole string, userID primitive.ObjectID, title string) error {
	storedTask, err := m.GetByTitle(title)
	if err != nil {
		return err
	}

	if storedTask.UserId != userID && userRole != "admin" {
		return errors.New("unauthorized to delete this task")
	}
	_, err = m.collection.DeleteOne(context.TODO(), bson.M{"title": title})
	if err != nil {
		return err
	}

	return nil
}

// func (m *UserDataManipulator) GetAllTasks(userRole string, userID primitive.ObjectID) ([]*domain.Task, error) {
// 	var tasks []*domain.Task

// 	filter := bson.M{"userid": userID}
// 	if userRole == "admin" {
// 		filter = bson.M{}
// 	}

// 	cursor, err := m.collection.Find(context.TODO(), filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.TODO())

// 	for cursor.Next(context.TODO()) {
// 		var task domain.Task
// 		if err = cursor.Decode(&task); err != nil {
// 			return nil, err
// 		}
// 		tasks = append(tasks, &task)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return tasks, nil
// }
