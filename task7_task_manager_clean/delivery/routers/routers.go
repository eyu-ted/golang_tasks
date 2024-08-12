package routers

import (
	"tskmgr/delivery/controllers"
	"tskmgr/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(usercontroller *controllers.Usercontroller, taskcontroller *controllers.TaskController) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", usercontroller.SignupController)
	router.POST("/login", usercontroller.LoginController)

	protected := router.Group("/")
	protected.Use(infrastructure.AuthMiddleware)
	{
		protected.POST("/task", taskcontroller.CreateTask)
		// get all tasks
		protected.GET("/tasks", taskcontroller.GetAllTasks)
		// get user tasks
		protected.GET("/mytasks", taskcontroller.GetUserTasks)
		// update user task
		protected.PUT("/task/:title", taskcontroller.UpdateTask)
		// get task by title
		protected.GET("/task/:title", taskcontroller.GetTaskByTitle)
		// delete task by title
		protected.DELETE("/task/:title", taskcontroller.DeleteTask)
	}

	return router
}
