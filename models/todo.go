package models

import (
	"fmt"
)

type TodoStatus string

type TodoResponse struct {
	ID          int32      `db:"id"`
	Description string     `db:"description"`
	Status      TodoStatus `db:"status"`
	UserID      int32      `db:"user_id"`
}

type Todo struct {
	ID          int32       `db:"id"`
	Description *string     `db:"description"`
	Status      *TodoStatus `db:"status"`
	UserID      *int32      `db:"user_id"`
}

type TodoService struct{}

func (c TodoService) GetAll() []TodoResponse {
	var todo []TodoResponse
	err := GetConnection().Select(&todo, "SELECT id,description,status,user_id FROM todo;")
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}

func (c TodoService) GetById(id int32) []TodoResponse {
	var todo []TodoResponse
	err := GetConnection().Select(&todo, fmt.Sprintf("SELECT id,description,status,user_id FROM todo WHERE id = %v", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}

func (c TodoService) GetByUserId(id int32) []TodoResponse {
	var todo []TodoResponse
	err := GetConnection().Select(&todo, fmt.Sprintf("SELECT id,description,status,user_id FROM todo WHERE user_id = %v", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}

func (c TodoService) Save(todo Todo) *TodoResponse {
	var insertedTodo TodoResponse
	query := fmt.Sprintf("INSERT INTO public.todo (description, status, user_id) VALUES ('%s', 'PENDING', %v) returning id, description, status, user_id", *todo.Description, todo.UserID)

	err := GetConnection().QueryRow(query).Scan(&insertedTodo.ID, &insertedTodo.Description, &insertedTodo.Status, &insertedTodo.UserID)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &insertedTodo
}

func (c TodoService) Update(todo Todo) *TodoResponse {
	var updatedTodo TodoResponse
	query := fmt.Sprintf("UPDATE public.todo SET %s WHERE id = %v returning id, description, status, user_id", makeUpdateQuery(todo), todo.ID)
	err := GetConnection().QueryRow(query).Scan(&updatedTodo.ID, &updatedTodo.Description, &updatedTodo.Status, &updatedTodo.UserID)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &updatedTodo
}
