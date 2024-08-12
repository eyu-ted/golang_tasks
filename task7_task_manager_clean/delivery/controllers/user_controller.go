package controllers

import (
	"net/http"
	"tskmgr/domain"

	"github.com/gin-gonic/gin"
)

type Usercontroller struct {
	MyuserUsecase domain.UserUsecaseInterface
}

func NewUsercontroller(usecase domain.UserUsecaseInterface) *Usercontroller {
	return &Usercontroller{
		MyuserUsecase: usecase,
	}
}

func (cont *Usercontroller) SignupController(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err := cont.MyuserUsecase.CreateUser(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func (cont *Usercontroller) LoginController(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	token, err := cont.MyuserUsecase.LogUser(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user logged in successfully", "token": token})
}
