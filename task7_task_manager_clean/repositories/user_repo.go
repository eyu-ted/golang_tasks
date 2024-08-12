package repositories

import (
	"context"
	"errors"
	"tskmgr/domain"
	"tskmgr/infrastructure"

	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDataManipulator struct {
	collection *mongo.Collection
}

func NewUserDataManipulator(collection *mongo.Collection) *UserDataManipulator {
	return &UserDataManipulator{
		collection: collection,
	}
}

func (m *UserDataManipulator) StoreUser(user *domain.User) error {
	user.UserID = primitive.NewObjectID()
	if _, err := m.GetByUsername(user.Username); err == nil {
		return errors.New("user already exists")
	}
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	_, err = m.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserDataManipulator) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := m.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserDataManipulator) LoginUser(user *domain.User) (string, error) {
	storedUser, err := m.GetByUsername(user.Username)
	if err != nil {
		return "", errors.New("user does not exist")
	}

	err = infrastructure.CheckPassword(storedUser.Password, user.Password)
	if err != nil {
		return "", err
	}
	claims := domain.Claims{
		UserId:    user.UserID,
		UserEmail: user.Email,
		Username:  user.Username,
		UserRole:  user.UserRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token, err := infrastructure.GetToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
