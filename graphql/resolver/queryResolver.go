package resolver

type QueryResolver struct{}

func (_ QueryResolver) Todo() *todoQueryResolver {
	return &todoQueryResolver{}
}
