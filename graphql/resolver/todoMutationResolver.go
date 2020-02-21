package resolver

import (
	"github.com/Davd96/graphql-go-todo/models"
)

type todoMutationResolver struct{}

type CreateUser struct {
	Input models.NewUser
}

func (_ *todoMutationResolver) CreateUser(input *CreateUser) *userResolver {
	data := userService.Save(models.NewUser{Name: "pruebauser"})
	return &userResolver{
		user: *data,
	}
}
