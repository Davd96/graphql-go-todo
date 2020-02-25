package models

type TodoStatus string

type TodoResponse struct {
	ID          int32      `db:"id"`
	Description string     `db:"description"`
	Status      TodoStatus `db:"status"`
	UserID      int32      `db:"user_id"`
}

type CreateTodo struct {
	Input struct {
		Description string
		UserID      int32
	}
}

type UpdateTodo struct {
	Input struct {
		ID          int32       `db:"id"`
		Description *string     `db:"description"`
		Status      *TodoStatus `db:"status"`
	}
}

type DeleteTodo struct {
	Input struct {
		ID int32
	}
}
