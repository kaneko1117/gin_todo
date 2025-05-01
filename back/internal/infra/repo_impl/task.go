package repo_impl

import (
	"context"
	"errors"
	"gin_todo/internal/entity/model"
	"gin_todo/internal/entity/query"

	"gorm.io/gorm"
)

type ITaskRepo interface {
	CreateTask(t model.Task) error
	GetTasks(userID int32) ([]*model.Task, error)
	ChangeTaskStatus(taskID int32, isChecked bool) error
}

type TaskRepo struct {
	db *gorm.DB
	q  *query.Query
	c  context.Context
}

func NewTaskRepo(db *gorm.DB) ITaskRepo {
	q := query.Use(db)
	c := context.Background()
	return &TaskRepo{
		db: db,
		q:  q,
		c:  c,
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
	var tasks []*model.Task
	err := r.db.Where("user_id = ?", userID).Order("id").Find(&tasks).Error
	if err != nil {
		return nil, errors.New("タスクの取得に失敗しました")
	}
	if len(tasks) == 0 {
		return nil, errors.New("タスクが存在しません")
	}
	return tasks, nil
}

func (r *TaskRepo) ChangeTaskStatus(taskID int32, isChecked bool) error {
	q := r.q
	c := r.c
	_ ,err:= q.WithContext(c).Task.Where(q.Task.ID.Eq(taskID)).Updates(map[string]interface{}{"IsChecked": isChecked})
	if err != nil {
		return errors.New("タスクの更新に失敗しました")
	}
	return nil
}
