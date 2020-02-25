package resolver

import (
	"github.com/Davd96/graphql-go-todo/models"
)

type todoMutationResolver struct{}

func (_ *todoMutationResolver) CreateTodo(todoInput models.CreateTodo) *todoResolver {
	data := todoService.Save(todoInput)
	return &todoResolver{
		todo: *data,
	}
}

func (_ *todoMutationResolver) UpdateTodo(todoInput models.UpdateTodo) *todoResolver {
	data := todoService.Update(todoInput)
	return &todoResolver{
		todo: *data,
	}
}

func (_ *todoMutationResolver) DeleteTodo(todoInput models.DeleteTodo) *todoResolver {

	data := todoService.Delete(todoInput)
	return &todoResolver{
		todo: *data,
	}
}

func (_ *todoMutationResolver) CreateUser(userInput models.CreateUser) *userResolver {
	data := userService.Save(userInput)
	return &userResolver{
		user: *data,
	}
}

func (_ *todoMutationResolver) UpdateUser(userInput models.UpdateUser) *userResolver {

	data := userService.Update(userInput)
	return &userResolver{
		user: *data,
	}
}

func (_ *todoMutationResolver) DeleteUser(userInput models.DeleteUser) *userResolver {

	data := userService.Delete(userInput)
	return &userResolver{
		user: *data,
	}
}
