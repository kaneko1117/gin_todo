package request

// graphqlではこのdtoは不要だけど、用意する
type CreateTaskRequest struct {
	UserID int32  `json:"user_id" validate:"required"`
	Tasks  string `json:"tasks" validate:"required"`
}
