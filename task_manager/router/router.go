package router

import (
    "github.com/gin-gonic/gin"
    "task_manager/controllers"
)

func SetupRouter(){
    rout := gin.Default()
    rout.GET("/tasks", controllers.GetAllTasks)
    rout.GET("/tasks/:id", controllers.GetTaskByID)
    rout.POST("/tasks", controllers.CreateTask)
    rout.PUT("/tasks/:id", controllers.UpdateTask)
    rout.DELETE("/tasks/:id", controllers.DeleteTask)
	rout.Run(":8080") 

}
