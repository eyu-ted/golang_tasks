package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {

    tasks := data.GetAllTasks()
    if len(tasks) ==0{
        c.IndentedJSON(http.StatusOK, gin.H{"message": "There is no any created task"})

    }
    c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
    id_param := c.Param("id")

    if id_param== ""{
        
        c.JSON(http.StatusBadRequest,gin.H{"error":"Task ID is required"})
        return
    }
    id, err := strconv.Atoi(id_param)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type task ID"})
        return
    }
    
    task, err := data.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.IndentedJSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
    var newTask models.Task
    if err := c.BindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    createdTask ,err:= data.CreateTask(newTask)
    if err != nil{
        c.JSON(http.StatusConflict ,gin.H{"error" : "Task with this ID already exists"})

    }else{
    c.JSON(http.StatusCreated, createdTask)}
}


func UpdateTask(c *gin.Context) {
    id_param := c.Param("id")
    if id_param== ""{

        c.JSON(http.StatusBadRequest,gin.H{"error":"Task ID is required"})
        return
    }

    id, err := strconv.Atoi(id_param)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
   
    var updatedTask models.Task
    if err := c.BindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    task, err := data.UpdateTask(id, updatedTask)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    if c.Param("id") == ""{

        c.JSON(http.StatusBadRequest,gin.H{"error":"Task ID is required"})
        return
    }

    err = data.DeleteTask(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"delete": "succesfully"})
}
