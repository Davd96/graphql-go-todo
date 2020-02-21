package resolver

type todoQueryResolver struct{}

type idFilter struct {
	Code int32
}

func (_ *todoQueryResolver) Todo(id idFilter) *todoResolver {
	data := todoService.GetById(id.Code)
	return &todoResolver{todo: data[0]}
}

func (_ *todoQueryResolver) User(id idFilter) *userResolver {
	data := userService.GetById(id.Code)
	return &userResolver{user: data[0]}
}

func (_ *todoQueryResolver) AllTodos() []*todoResolver {
	data := todoService.GetAll()
	list := make([]*todoResolver, int(len(data)))

	for i := range list {
		list[i] = &todoResolver{
			todo: data[i],
		}
	}

	return list
}

func (_ *todoQueryResolver) AllUsers() []*userResolver {
	data := userService.GetAll()

	list := make([]*userResolver, len(data))

	for i := range list {
		list[i] = &userResolver{
			user: data[i],
		}
	}

	return list
}
