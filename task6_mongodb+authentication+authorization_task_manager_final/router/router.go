package router

import (
    "github.com/gin-gonic/gin"
    "task_manager/controllers"
    "task_manager/middleware"

)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/register", controllers.RegisterUser)
    r.POST("/login", controllers.LoginUser)

    protected := r.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {   
        protected.GET("/tasks", controllers.GetAllTasks)
        protected.GET("/tasks/:id", controllers.GetTaskByID)
        protected.POST("/tasks", controllers.CreateTask)
        protected.PUT("/tasks/:id", controllers.UpdateTaskUsers)
        protected.DELETE("/tasks/:id", controllers.DeleteTask)
        
        
    }

    admin := r.Group("/admin")
    admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
    {   
        admin.GET("/tasks", controllers.GetAllTasks)
        admin.GET("/tasks/:id", controllers.GetTaskByID)
        admin.PUT("/tasks/:id", controllers.UpdateTaskUsers)
        admin.DELETE("/tasks/:id", controllers.DeleteTask)
    }

    return r
}
