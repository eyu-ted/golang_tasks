package usecases

import (
	// "errors"
	// "time"
	"tskmgr/domain"
	// "tskmgr/infrastructure"
	// "github.com/golang-jwt/jwt"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	MyUserRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		MyUserRepo: repo,
	}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.MyUserRepo.StoreUser(user)
}

func (u *UserUsecase) LogUser(user *domain.User) (string, error) {

	return u.MyUserRepo.LoginUser(user)
}
