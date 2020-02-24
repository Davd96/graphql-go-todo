package resolver

import (
	"strconv"

	"github.com/Davd96/graphql-go-todo/models"
	graphql "github.com/guzmanweb/graphql-go"
)

type userResolver struct {
	user models.UserResponse
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.user.ID)))
}

func (r *userResolver) Name() string {
	return r.user.Name
}

func (r *userResolver) Todos() []*todoResolver {
	data := todoService.GetByUserId(r.user.ID)

	list := make([]*todoResolver, len(data))

	for i := range list {
		list[i] = &todoResolver{
			todo: data[i],
		}
	}

	return list
}
