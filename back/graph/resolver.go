//go:generate go run github.com/99designs/gqlgen generate
package graph

import "gin_todo/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TaskUseCase usecase.ITaskUseCase
	UserUseCase usecase.IUserUseCase
}

func NewResolver(taskUseCase usecase.ITaskUseCase, userUseCase usecase.IUserUseCase) *Resolver {
	return &Resolver{
		TaskUseCase: taskUseCase,
		UserUseCase: userUseCase,
	}
}
