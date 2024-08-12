package controllers

import (
	"errors"
	"net/http"
	"tskmgr/domain"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	MyTaskUsecase domain.TaskUsecaseInterface
}

func NewTaskController(usecase domain.TaskUsecaseInterface) *TaskController {
	return &TaskController{
		MyTaskUsecase: usecase,
	}
}

func getclaim(c *gin.Context) (*domain.Claims, error) {
	claim, exists := c.Get("claim")
	if !exists {
		return nil, errors.New("claim not set")
	}

	userClaims, ok := claim.(domain.Claims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return &userClaims, nil
}

func (cont *TaskController) CreateTask(c *gin.Context) {
	claim, err := getclaim(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newTask := &domain.Task{}
	if err := c.ShouldBindJSON(newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	userid := claim.UserId
	newTask.UserId = userid

	// claims := c.MustGet("claim").(domain.Claims)

	newTask, err = cont.MyTaskUsecase.CreateTask(newTask)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

func (cont *TaskController) GetTaskByTitle(c *gin.Context) {
	title := c.Param("title")
	task, err := cont.MyTaskUsecase.GetTaskByTitle(title)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, task)
}

// get all tasks
func (cont *TaskController) GetAllTasks(c *gin.Context) {
	claim, err := getclaim(c)
	// title := c.Param("title")
	// newTask, err = cont.MyTaskUsecase.UpdateTask()
	tasks, err := cont.MyTaskUsecase.GetAllTasks(claim.UserRole, claim.UserId)
	// tasks, err := cont.MyTaskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, tasks)
}

// get user tasks
func (cont *TaskController) GetUserTasks(c *gin.Context) {
	claim, err := getclaim(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userid := claim.UserId
	tasks, err := cont.MyTaskUsecase.GetUserTasks(userid)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, tasks)
}

// update user task if user id is the same or user is admin
func (cont *TaskController) UpdateTask(c *gin.Context) {
	claim, err := getclaim(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	title := c.Param("title")

	newTask := &domain.Task{}
	if err := c.ShouldBindJSON(newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	newTask, err = cont.MyTaskUsecase.UpdateTask(claim.UserRole, claim.UserId, title, newTask)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

// delete task by title
func (cont *TaskController) DeleteTask(c *gin.Context) {
	claim, err := getclaim(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	title := c.Param("title")

	err = cont.MyTaskUsecase.DeleteTask(claim.UserRole, claim.UserId, title)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "task deleted"})
}
