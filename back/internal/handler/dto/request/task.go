package request

// graphqlではこのdtoは不要だけど、用意する
type CreateTaskRequest struct {
	UserID int32  `json:"user_id" validate:"required"`
	Tasks  string `json:"tasks" validate:"required"`
}

type ChangeTaskStatusRequest struct {
	TaskID   int32 `json:"task_id" validate:"required"`
	IsChecked bool  `json:"is_checked" validate:"required,boolean"`
}
