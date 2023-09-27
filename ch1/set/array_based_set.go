package set

import (
	"slices"
	"fmt"
)

type ArrayBasedSet[T comparable] struct {
	// unsorted array
	array []T
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
	if s.Size() != t.Size() {
		return false
	}

	var (
		hash = make(map[T]struct{})
	)

	for _, v := range s.array {
		hash[v] = struct{}{}
	}

	for i := 0; i < t.Size(); i++ {
		if _, exist := hash[t.Members()[i]]; !exist {
			return false
		}
	}
	return true
}

func (s *ArrayBasedSet[T]) Get(index int) (value *T) {
	if index < 0 || index >= len(s.array) {
		return nil
	}
	return &s.array[index]
}

func (s *ArrayBasedSet[T]) Intersect(t Set[T]) Set[T] {
	var (
		hash = make(map[T]struct{})
		result []T
		resultSet Set[T]
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

	return resultSet
}

func (s *ArrayBasedSet[T]) IsNull() bool {
	return len(s.array) == 0
}

func (s *ArrayBasedSet[T]) IsSubsetOf(t Set[T]) bool {
	var (
		hash = make(map[T]struct{})
	)

	for _, v := range t.Members() {
		hash[v] = struct{}{}
	}

	for _, v := range s.array {
		if _, exist := hash[v]; !exist {
			return false
		}
	}

	return true
}

func (s *ArrayBasedSet[T]) Members() []T {
	return s.array
}

func (s *ArrayBasedSet[T]) Remove(value T) {
	index := s.Search(value)
	s.array = slices.Delete(s.array, index, index + 1)
}

// If no value is found, return -1. Othewise, return the index.
func (s *ArrayBasedSet[T]) Search(value T) (index int) {
	var (
		result = -1
	)

	for i := 0; i < len(s.array); i++ {
		if s.array[i] == value {
			return i
		}
	}

	return result
}

func (s *ArrayBasedSet[T]) Size() int {
	return len(s.array)
}

func (s *ArrayBasedSet[T]) Union(t Set[T]) Set[T] {
	var (
		hash = make(map[T]struct{})
		resultSet Set[T]
	)

	for _, v := range s.array {
		hash[v] = struct{}{}
	}

	for _, v := range t.Members() {
		delete(hash, v)
	}

	for key, _ := range hash {
		resultSet.Add(key)
	}

	return resultSet
}