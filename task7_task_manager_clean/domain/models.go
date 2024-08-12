package domain

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID   primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username" binding:"required"`
	Email    string             `bson:"email" binding:"required"`
	Password string             `bson:"password" binding:"required"`
	UserRole string             `bson:"role" binding:"required"`
}

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    primitive.ObjectID `bson:"userid"`
	OwnerID   primitive.ObjectID `bson:"ownerid"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
}

type Claims struct {
	UserId    primitive.ObjectID `json:"userid"`
	UserEmail string             `json:"useremail"`
	Username  string             `json:"username"`
	UserRole  string             `json:"role"`
	jwt.StandardClaims
}
