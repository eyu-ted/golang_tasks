package usecases

import (
	// "errors"
	"tskmgr/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	MyTaskRepo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		MyTaskRepo: repo,
	}
}

func (u *TaskUsecase) CreateTask(task *domain.Task) (*domain.Task, error) {

	return u.MyTaskRepo.StoreTask(task)
}

func (u *TaskUsecase) GetTaskByTitle(title string) (*domain.Task, error) {
	return u.MyTaskRepo.GetByTitle(title)
}

// func (u *TaskUsecase) GetAllTasks() ([]*domain.Task, error) {
// 	return u.MyTaskRepo.GetAllTasks()
// }

func (u *TaskUsecase) GetUserTasks(userid primitive.ObjectID) ([]*domain.Task, error) {
	return u.MyTaskRepo.GetUserTasks(userid)
}

func (u *TaskUsecase) UpdateTask(userRole string, userID primitive.ObjectID, title string, task *domain.Task) (*domain.Task, error) {
	return u.MyTaskRepo.UpdateTask(userRole, userID, title, task)
}

func (u *TaskUsecase) DeleteTask(userRole string, userID primitive.ObjectID, title string) error {

	return u.MyTaskRepo.DeleteTask(userRole, userID, title)
}

func (u *TaskUsecase) GetAllTasks(userRole string, userID primitive.ObjectID) ([]*domain.Task, error) {
	return u.MyTaskRepo.GetAllTasks(userRole, userID)
}
