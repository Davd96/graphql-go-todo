package services

import (
	"fmt"

	"github.com/Davd96/graphql-go-todo/models"
)

type TodoService struct{}

func (c TodoService) GetAll() []models.TodoResponse {
	var todo []models.TodoResponse
	err := GetConnection().Select(&todo, "SELECT id,description,status,user_id FROM todo;")
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}

func (c TodoService) GetById(id int32) []models.TodoResponse {
	var todo []models.TodoResponse
	err := GetConnection().Select(&todo, fmt.Sprintf("SELECT id,description,status,user_id FROM todo WHERE id = %v", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}

func (c TodoService) GetByUserId(id int32) []models.TodoResponse {
	var todo []models.TodoResponse
	err := GetConnection().Select(&todo, fmt.Sprintf("SELECT id,description,status,user_id FROM todo WHERE user_id = %v", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}

func (c TodoService) Save(todo models.CreateTodo) *models.TodoResponse {
	var insertedTodo models.TodoResponse
	query := fmt.Sprintf("INSERT INTO public.todo (description, status, user_id) VALUES ('%s', 'PENDING', %v) returning id, description, status, user_id", todo.Input.Description, todo.Input.UserID)

	err := GetConnection().QueryRow(query).Scan(&insertedTodo.ID, &insertedTodo.Description, &insertedTodo.Status, &insertedTodo.UserID)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &insertedTodo
}

func (c TodoService) Update(todo models.UpdateTodo) *models.TodoResponse {
	var updatedTodo models.TodoResponse
	query := fmt.Sprintf("UPDATE public.todo SET %s WHERE id = %v returning id, description, status, user_id", makeUpdateQuery(todo.Input), todo.Input.ID)
	err := GetConnection().QueryRow(query).Scan(&updatedTodo.ID, &updatedTodo.Description, &updatedTodo.Status, &updatedTodo.UserID)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &updatedTodo
}

func (c TodoService) Delete(todo models.DeleteTodo) *models.TodoResponse {
	var deletedTodo models.TodoResponse

	query := fmt.Sprintf("DELETE FROM public.todo WHERE id=%v returning id, description, status, user_id", todo.Input.ID)

	err := GetConnection().QueryRow(query).Scan(&deletedTodo.ID, &deletedTodo.Description, &deletedTodo.Status, &deletedTodo.UserID)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &deletedTodo
}
