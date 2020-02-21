package resolver

type MutationResolver struct{}

func (_ MutationResolver) Todo() *todoMutationResolver {
	return &todoMutationResolver{}
}
