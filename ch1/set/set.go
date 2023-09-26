package set

type Set interface {
	Add(value any) error
	Equal(s Set) bool
	Get(index int) (value any)
	Init(values ...any) Set
	Intersect(s Set) Set
	IsNull() bool
	IsSubsetOf(s Set) bool
	Remove(value any)
	Search(value any) (index int)
	Size() int
	Union(s Set) Set
}
