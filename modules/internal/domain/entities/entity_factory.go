package entities

type (
	Factory[T any] interface {
		New() T
	}
	ElementFactory[T any] struct{}
)

func (f ElementFactory[T]) New() T {
	var entity T
	return &entity
}
