package controllers

import (
	"fmt"
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"task_manager/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)
func GetAllTasks(c *gin.Context) {

    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
        return
    }

    claims, ok := user.(*middleware.Claims)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims"})
        return
    }

 
    tasks, err := data.GetAllTasks(claims.UserID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tasks)
}
func GetTaskByID(c *gin.Context) {
    id := c.Param("id")
    taskId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
        return
    }

    claims, ok := user.(*middleware.Claims)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims"})
        return
    }

    task, err := data.GetTaskByID(taskId.Hex(), claims.UserID)
    if err != nil {
        if err.Error() == "task not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, task)
}



func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		return
	}

	claims, ok := user.(*middleware.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims"})
		return
	}

	createdTask, err := data.CreateTask(claims, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdTask)
}

func UpdateTaskUsers(c *gin.Context) {
	id := c.Param("id")
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask.ID = taskId

	user, exists := c.Get("user")
	if !exists {
		fmt.Println("user1", user)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		return
	}
	fmt.Println("user2", user)

	claims, ok := user.(*middleware.Claims)
	if !ok {
		fmt.Println("claims1", claims)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims"})
		return
	}
	fmt.Println("claims2", claims)

	task, err := data.UpdateTaskUsers(claims, updatedTask)
	if err != nil {
		fmt.Println("task1", task)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("task2", task)

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    taskId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    // Extract user from context (set by AuthMiddleware)
    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
        return
    }

    claims, ok := user.(*middleware.Claims)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims"})
        return
    }

    // Call the data layer to delete the task, ensuring the user is the owner
    err = data.DeleteTask(taskId.Hex(), claims.UserID)
    if err != nil {
        if err.Error() == "task not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        } else if err.Error() == "not authorized to delete this task" {
            c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this task"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}


// func DeleteAllTasks(c *gin.Context) {
// 	if err := data.DeleteAllTasks(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "All tasks deleted successfully"})
// }




