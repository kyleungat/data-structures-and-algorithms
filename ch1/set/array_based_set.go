package set

import (
	"slices"
	"fmt"
)

type ArrayBasedSet[T comparable] struct {
	// unsorted array
	array []T
}

func NewArrayBasedSet[T comparable]() *ArrayBasedSet[T] {
	return &ArrayBasedSet[T]{}
}

func (s *ArrayBasedSet[T]) Add(value T) error {
	exist := slices.Contains(s.array, value)
	if exist {
		return fmt.Errorf("element %v already exists", value)
	}
	s.array = append(s.array, value)
	return nil
}

func (s *ArrayBasedSet[T]) Equal(t Set[T]) bool {
	intersection := s.Intersect(t) 
	return intersection.Size() == s.Size() && s.Size() == t.Size()
}

func (s *ArrayBasedSet[T]) Intersect(t Set[T]) Set[T] {
	var (
		hash = make(map[T]struct{})
		result []T
		resultSet = ArrayBasedSet[T]{}
	)

	for _, v := range s.array {
		hash[v] = struct{}{}
	}

	for _, v := range t.Members() {
		if _, exist := hash[v]; exist {
			result = append(result, v)
		}
	}

	for _, v := range result {
		resultSet.Add(v)
	}

	return &resultSet
}

func (s *ArrayBasedSet[T]) IsNull() bool {
	return len(s.array) == 0
}

func (s *ArrayBasedSet[T]) IsSubsetOf(t Set[T]) bool {
	intersection := s.Intersect(t) 
	return intersection.Size() == s.Size() 
}

func (s *ArrayBasedSet[T]) Members() []T {
	return s.array
}

func (s *ArrayBasedSet[T]) Remove(value T) {
	var (
		index = -1
	)

	for i := 0; i < len(s.array); i++ {
		if s.array[i] == value {
			index = i
		}
	}

	if index != -1 {
		s.array = slices.Delete(s.array, index, index + 1)
	}
}

func (s *ArrayBasedSet[T]) Search(value T) bool {
	for i := 0; i < len(s.array); i++ {
		if s.array[i] == value {
			return true
		}
	}

	return false
}

func (s *ArrayBasedSet[T]) Size() int {
	return len(s.array)
}

func (s *ArrayBasedSet[T]) Union(t Set[T]) Set[T] {
	var (
		// hash = make(map[T]struct{})
		resultSet = ArrayBasedSet[T]{}
	)

	for _, v := range s.array {
		resultSet.Add(v)
	}

	for _, v := range t.Members() {
		resultSet.Add(v)
	}
	
	return &resultSet
}