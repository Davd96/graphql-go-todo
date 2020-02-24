package resolver

import (
	"strconv"

	"github.com/Davd96/graphql-go-todo/models"
	graphql "github.com/guzmanweb/graphql-go"
)

type todoResolver struct {
	todo models.TodoResponse
}

func (r *todoResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.todo.ID)))
}

func (r *todoResolver) Description() string {
	return r.todo.Description
}

func (r *todoResolver) Status() models.TodoStatus {
	return r.todo.Status
}

func (r *todoResolver) User() *userResolver {

	data := userService.GetById(r.todo.UserID)

	return &userResolver{
		user: data[0],
	}
}
