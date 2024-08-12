package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "tskmgr/domain"
)

type UserRepository interface {
	StoreUser(user *User) error
	GetByUsername(username string) (*User, error)
	LoginUser(user *User) (string, error)
}

type TaskRepository interface {
	GetAllTasks(userRole string, userID primitive.ObjectID) ([]*Task, error)
	StoreTask(task *Task) (*Task, error)
	GetByTitle(title string) (*Task, error)
	// GetAllTasks() ([]*Task, error)
	GetUserTasks(userid primitive.ObjectID) ([]*Task, error)
	UpdateTask(userRole string, userID primitive.ObjectID, title string, task *Task) (*Task, error)
	DeleteTask(userRole string, userID primitive.ObjectID, title string) error
}

type UserUsecaseInterface interface {
	CreateUser(user *User) error
	LogUser(user *User) (string, error)
}

type TaskUsecaseInterface interface {
	GetAllTasks(userRole string, userID primitive.ObjectID) ([]*Task, error)
	CreateTask(task *Task) (*Task, error)
	GetTaskByTitle(title string) (*Task, error)
	// GetAllTasks() ([]*Task, error)
	GetUserTasks(userid primitive.ObjectID) ([]*Task, error)
	UpdateTask(userRole string, userID primitive.ObjectID, title string, task *Task) (*Task, error)
	DeleteTask(userRole string, userID primitive.ObjectID, title string) error
}
