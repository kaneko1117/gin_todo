package repo_impl

import (
	"errors"
	"gin_todo/internal/infra/model"
	"gin_todo/internal/infra/query"

	"gorm.io/gorm"
)





type ITaskRepo interface {
	CreateTask(t model.Task) (error)
}

type TaskRepo struct {
	db *gorm.DB
	q *query.Query
}

func NewTaskRepo(db *gorm.DB) ITaskRepo {
	q := query.Use(db)
	return &TaskRepo{
		db: db,
		q:  q,
	}
}

func (r *TaskRepo) CreateTask(t model.Task) (error) {
	q := r.q
	err := q.Task.Create(&t)
	if err != nil {
		return errors.New("タスクの登録に失敗しました")
	}
	return nil
}
