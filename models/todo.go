package models

import "fmt"

type Todo struct {
	ID          int32  `db:"id"`
	Description string `db:"description"`
	Status      bool   `db:"status"`
	UserID      int32  `db:"user_id"`
}

type TodoService struct{}

func (c TodoService) GetAll() []Todo {
	var todo []Todo
	err := GetConnection().Select(&todo, "SELECT id,description,status,user_id FROM todo;")
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}

func (c TodoService) GetById(id int32) []Todo {
	var todo []Todo
	err := GetConnection().Select(&todo, fmt.Sprintf("SELECT id,description,status,user_id FROM todo WHERE id = %v", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}

func (c TodoService) GetByUserId(id int32) []Todo {
	var todo []Todo
	err := GetConnection().Select(&todo, fmt.Sprintf("SELECT id,description,status,user_id FROM todo WHERE user_id = %v", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}
