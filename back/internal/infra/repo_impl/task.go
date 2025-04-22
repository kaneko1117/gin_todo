package repo_impl

import (
	"errors"
	"gin_todo/internal/entity/model"
	"gin_todo/internal/entity/query"

	"gorm.io/gorm"
)

type ITaskRepo interface {
	CreateTask(t model.Task) error
	GetTasks(userID int32) ([]*model.Task, error)
}

type TaskRepo struct {
	db *gorm.DB
	q  *query.Query
}

func NewTaskRepo(db *gorm.DB) ITaskRepo {
	q := query.Use(db)
	return &TaskRepo{
		db: db,
		q:  q,
	}
}

func (r *TaskRepo) CreateTask(t model.Task) error {
	q := r.q
	err := q.Task.Create(&t)
	if err != nil {
		return errors.New("タスクの登録に失敗しました")
	}
	return nil
}

func (r *TaskRepo) GetTasks(userID int32) ([]*model.Task, error) {
	q := r.q
	tasks, err := q.Task.Where(query.Task.UserID.Eq(userID)).Find()
	if err != nil {
		return nil, errors.New("タスクの取得に失敗しました")
	}
	if len(tasks) == 0 {
		return nil, errors.New("タスクが存在しません")
	}
	return tasks, nil
}
