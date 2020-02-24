package resolver

import (
	"github.com/Davd96/graphql-go-todo/models"
)

type todoMutationResolver struct{}

type CreateTodo struct {
	Input models.Todo
}

type UpdateTodo struct {
	Input models.Todo
}

type CreateUser struct {
	Input models.User
}

type UpdateUser struct {
	Input models.User
}

func (_ *todoMutationResolver) CreateTodo(todoInput *CreateTodo) *todoResolver {
	data := todoService.Save(todoInput.Input)
	return &todoResolver{
		todo: *data,
	}
}

func (_ *todoMutationResolver) UpdateTodo(todoInput *UpdateTodo) *todoResolver {
	data := todoService.Update(todoInput.Input)
	return &todoResolver{
		todo: *data,
	}
}

func (_ *todoMutationResolver) CreateUser(userInput *CreateUser) *userResolver {
	data := userService.Save(userInput.Input)
	return &userResolver{
		user: *data,
	}
}

func (_ *todoMutationResolver) UpdateUser(userInput *UpdateUser) *userResolver {

	data := userService.Update(userInput.Input)
	return &userResolver{
		user: *data,
	}
}
