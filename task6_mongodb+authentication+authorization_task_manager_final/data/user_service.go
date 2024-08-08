package data

import (
	"context"
	"time"

	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection

func InitUserCollection() {
	UserCollection = Client.Database("task_manager").Collection("users")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.Password, _ = HashPassword(user.Password)
	return UserCollection.InsertOne(ctx, user)
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return &user, err
}
