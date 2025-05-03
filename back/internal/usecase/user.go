package usecase

import (
	"gin_todo/internal/infra/repo_impl"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IUserUseCase interface {
	Login(userName string, password string) (string, error)
}

type UserUseCase struct {
	userRepository repo_impl.IUserRepo
}

func NewUserUseCase(userRepository repo_impl.IUserRepo) IUserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}
func (u *UserUseCase) Login(userName string, password string) (string, error) {
	err := u.userRepository.Login(userName, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userName,
		"registeredClaims": jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString, nil
}
