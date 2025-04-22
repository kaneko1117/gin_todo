package usecase

import (
	"gin_todo/internal/entity/model"
	"gin_todo/internal/infra/repo_impl"
	"time"
)

type ITaskUseCase interface {
	CreateTask(userID int32, tasks string) error
	GetTasks(userID int32) ([]*model.Task, error)
}

type TaskUseCase struct {
	taskRepository repo_impl.ITaskRepo
}

func NewTaskUseCase(taskRepository repo_impl.ITaskRepo) ITaskUseCase {
	return &TaskUseCase{
		taskRepository: taskRepository,
	}
}
func (u *TaskUseCase) CreateTask(userID int32, tasks string) error {
	task := model.Task{
		UserID:    userID,
		Tasks:     tasks,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	err := u.taskRepository.CreateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (u *TaskUseCase) GetTasks(userID int32) ([]*model.Task, error) {
	tasks, err := u.taskRepository.GetTasks(userID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
