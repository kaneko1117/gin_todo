package repo_impl

import (
	"context"
	"errors"
	"gin_todo/internal/entity/query"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepo interface {
	Login(userName string, password string) error
}

type UserRepo struct {
	db *gorm.DB
	q  *query.Query
	c  context.Context
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	q := query.Use(db)
	c := context.Background()
	return &UserRepo{
		db: db,
		q:  q,
		c:  c,
	}
}

func (r *UserRepo) Login(userName string, password string) error {
	q := r.q
	c := r.c
	user, err := q.WithContext(c).User.Where(q.User.UserName.Eq(userName)).First()
	if err != nil {
		return errors.New("ユーザー名またはパスワードが間違っています")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("ユーザー名またはパスワードが間違っています")
	}

	return nil
}
