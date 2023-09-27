package set

type Set[T comparable] interface {
	Add(value T) error
	Equal(s Set[T]) bool
	Intersect(s Set[T]) Set[T]
	IsNull() bool
	IsSubsetOf(s Set[T]) bool
	Members() []T
	Remove(value T)
	Search(value T) bool
	Size() int
	Union(s Set[T]) Set[T]
}
