package set

import (
	"fmt"
)

type HashBasedSet[T comparable] struct {
	hash map[T]struct{}
}

func NewHashBasedSet[T comparable]() *HashBasedSet[T] {
	return &HashBasedSet[T]{
		hash: make(map[T]struct{}),
	}
}

func (s *HashBasedSet[T]) Add(value T) error {
	if _, exist := s.hash[value]; exist {
		return fmt.Errorf("element %v already exists", value)
	}
	s.hash[value] = struct{}{}
	return nil
}

func (s *HashBasedSet[T]) Equal(t Set[T]) bool {
	intersection := s.Intersect(t) 
	return intersection.Size() == s.Size() && s.Size() == t.Size()
}

func (s *HashBasedSet[T]) Intersect(t Set[T]) Set[T] {
	var (
		result []T
		resultSet = HashBasedSet[T]{
			hash: make(map[T]struct{}),
		}
	)

	for _, v := range t.Members() {
		if _, exist := s.hash[v]; exist {
			result = append(result, v)
		}
	}

	for _, v := range result {
		resultSet.Add(v)
	}

	return &resultSet
}

func (s *HashBasedSet[T]) IsNull() bool {
	return len(s.hash) == 0
}

func (s *HashBasedSet[T]) IsSubsetOf(t Set[T]) bool {
	intersection := s.Intersect(t) 
	return intersection.Size() == s.Size() 
}

func (s *HashBasedSet[T]) Members() []T {
	var (
		members []T
	)

	for key, _ := range s.hash {
		members = append(members, key)
	}

	return members
}

func (s *HashBasedSet[T]) Remove(value T) {
	delete(s.hash, value)
}

func (s *HashBasedSet[T]) Search(value T) bool {
	if _, exist := s.hash[value]; exist {
		return true
	}
	return false
}

func (s *HashBasedSet[T]) Size() int {
	return len(s.hash)
}

func (s *HashBasedSet[T]) Union(t Set[T]) Set[T] {
	var (
		resultSet = HashBasedSet[T]{
			hash: make(map[T]struct{}),
		}
	)

	for key, _ := range s.hash {
		_ = resultSet.Add(key)
	}

	for _, value := range t.Members() {
		_ = resultSet.Add(value)
	}

	return &resultSet
}