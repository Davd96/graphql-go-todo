package resolver

import (
	"strconv"

	graphql "github.com/Davd96/graphql-go"
	"github.com/Davd96/graphql-go-todo/models"
)

type todoResolver struct {
	todo models.Todo
}

func (r *todoResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.todo.ID)))
}

func (r *todoResolver) Text() string {
	return r.todo.Description
}

func (r *todoResolver) Status() bool {
	return r.todo.Status
}

func (r *todoResolver) User() *userResolver {

	data := userService.GetById(r.todo.UserID)

	return &userResolver{
		user: data[0],
	}
}
