package data

import (
	"errors"
	"task_manager/models"

)

var tasks []models.Task

func GetAllTasks() []models.Task {
    return tasks
}

func GetTaskByID(id int) (*models.Task, error) {

    for _, task := range tasks {
        if task.ID == id {
            return &task, nil
        }
    }
    return nil, errors.New("task not found")
}

func CreateTask(task models.Task) (models.Task ,error){
    for _, val := range tasks {
        if task.ID == val.ID  {
            return task , errors.New("the ID already exists. the ID MUST be UNIQE")
        }
    }
    tasks = append(tasks, task)
    return task ,nil
}

func UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i] = updatedTask
            return &tasks[i], nil
        }
    }
    return nil, errors.New("task not found")
}

func DeleteTask(id int) error {
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return nil
        }
    }
    return errors.New("task not found")
}









