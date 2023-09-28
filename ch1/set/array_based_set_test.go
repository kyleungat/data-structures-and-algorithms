package set

import (
	"testing"
	"reflect"
	"fmt"
	"slices"
)

func TestArrayBasedSetAdd(t *testing.T) {
	type Want struct {
		members []int
		err error 
	}
	tests := []struct {
		name    string
		initSet func() Set[int]
		add 	int
		want    Want
	}{
		{
			"Given an empty set, then run Add(1), {1} is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				return set
			},
			1,
			Want{
				[]int{1},
				nil,
			},
		},
		{
			"Given a set {1, 2, 3}, then run Add(2), err is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				return set
			},
			2,
			Want{
				[]int{1,2,3},
				fmt.Errorf("element %v already exists", 2),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := tt.initSet()
			_ = set.Add(tt.add)

			members := set.Members()
			slices.Sort(members)
			if !reflect.DeepEqual(tt.want.members, members) {
				t.Fatalf("Add() results in %v, want match for %v", members, tt.want.members)
			}

		})
	}

}

func TestArrayBasedSetEqual(t *testing.T) {
	tests := []struct {
		name    string
		initSets func() (Set[int], Set[int])
		want    bool
	}{
		{
			"Given {}, {}, then run Equal(), true is expected",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set2 := NewArrayBasedSet[int]()
				return set1, set2
			},
			true,
		},
		{
			"Given {}, {1}, then run Equal(), false is expected",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				return set1, set2
			},
			false,
		},
		{
			"Given {1,2,3}, {1,2,3}, then run Equal(), true is expected",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(1)
				set1.Add(2)
				set1.Add(3)
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				set2.Add(2)
				set2.Add(3)
				return set1, set2
			},
			true,
		},
		{
			"Given {1,2}, {1,2,3}, then run Equal(), false is expected",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(1)
				set1.Add(2)
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				set2.Add(2)
				set2.Add(3)
				return set1, set2
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set1, set2 := tt.initSets()
			equal := set1.Equal(set2)

			if tt.want != equal {
				t.Fatalf("Equal() = %v, want match for %v", equal, tt.want)
			}
		})
	}
}

func TestArrayBasedSetIntersect(t *testing.T) {
	tests := []struct {
		name    string
		initSets func() (Set[int], Set[int])
		want    []int
	}{
		{
			"{} + {} = {}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set2 := NewArrayBasedSet[int]()
				return set1, set2
			},
			nil,
		},
		{
			"{} + {1,2} = {}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				set2.Add(2)
				return set1, set2
			},
			nil,
		},
		{
			"{1,2,3} + {1,2} = {1,2}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(1)
				set1.Add(2)
				set1.Add(3)
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				set2.Add(2)
				return set1, set2
			},
			[]int{1,2},
		},
		{
			"{1,2,3} + {4,5,6} = {}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(1)
				set1.Add(2)
				set1.Add(3)
				set2 := NewArrayBasedSet[int]()
				set2.Add(4)
				set2.Add(5)
				set2.Add(6)
				return set1, set2
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set1, set2 := tt.initSets()
			set := set1.Intersect(set2)

			members := set.Members()
			slices.Sort(members)
			if !reflect.DeepEqual(tt.want, members) {
				t.Fatalf("Intersect() = %v, want match for %v", members, tt.want)
			}
		})
	}

}

func TestArrayBasedSetIsNull(t *testing.T) {
	tests := []struct {
		name    string
		initSet func() Set[int]
		want    bool
	}{
		{
			"Given an empty set, then run isNull(), true is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				return set
			},
			true,
		},
		{
			"Given a non-empty set, then run isNull(), false is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				set.Add(1)
				return set
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := tt.initSet()
			isNull := set.IsNull()

			if tt.want != isNull {
				t.Fatalf("IsNull() = %v, want match for %v", isNull, tt.want)
			}
		})
	}

}

func TestArrayBasedSetIsSubsetOf(t *testing.T) {
	tests := []struct {
		name    string
		initSets func() (Set[int], Set[int])
		want    bool
	}{
		{
			"{} is subset of {}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set2 := NewArrayBasedSet[int]()
				return set1, set2
			},
			true,
		},
		{
			"{} is subset of {1}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				return set1, set2
			},
			true,
		},
		{
			"{1} is not subset of {}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(1)
				set2 := NewArrayBasedSet[int]()
				return set1, set2
			},
			false,
		},
		{
			"{1} is subset of {1}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(1)
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				return set1, set2
			},
			true,
		},
		{
			"{1} is subset of {1,3,5}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(1)
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				set2.Add(3)
				set2.Add(5)
				return set1, set2
			},
			true,
		},
		{
			"{2,3} is not subset of {1,3,5}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(2)
				set1.Add(3)
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				set2.Add(3)
				set2.Add(5)
				return set1, set2
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set1, set2 := tt.initSets()
			isSubset := set1.IsSubsetOf(set2)

			if tt.want != isSubset {
				t.Fatalf("%v's IsSubsetOf(%v) = %v, want match for %v", set1, set2, isSubset, tt.want)
			}
		})
	}

}

func TestArrayBasedSetMembers(t *testing.T) {
	tests := []struct {
		name    string
		initSet func() Set[int]
		want    []int
	}{
		{
			"Given an empty set, then run Members(), return []",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				return set
			},
			nil,
		},
		{
			"Given an {1,2,3}, then run Members(), return [1,2,3]",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				return set
			},
			[]int{1,2,3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := tt.initSet()
			members := set.Members()

			slices.Sort(members)
			if !reflect.DeepEqual(tt.want, members) {
				t.Fatalf("Members() = %v, want match for %v", members, tt.want)
			}
		})
	}

}

func TestArrayBasedSetRemove(t *testing.T) {
	tests := []struct {
		name    string
		initSet func() Set[int]
		remove int
		want    []int
	}{
		{
			"Given an empty set, then run Remove(1), an empty set is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				return set
			},
			1,
			nil,
		},
		{
			"Given a set {1,2}, then run Remove(2), {1} is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				set.Add(1)
				set.Add(2)
				return set
			},
			2,
			[]int{1},
		},
		{
			"Given a set {1,2,3}, then run Remove(1), {2,3} is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				return set
			},
			1,
			[]int{2,3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := tt.initSet()
			set.Remove(tt.remove)

			members := set.Members()
			slices.Sort(members)
			if !reflect.DeepEqual(tt.want, members) {
				t.Fatalf("Remove() = %v, want match for %v", members, tt.want)
			}
		})
	}

}

func TestArrayBasedSetSearch(t *testing.T) {
	tests := []struct {
		name    string
		initSet func() Set[int]
		search int
		want    bool
	}{
		{
			"Given an empty set, then run Search(1), false is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				return set
			},
			1,
			false,
		},
		{
			"Given a set {1,2}, then run Search(2), true is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				set.Add(1)
				set.Add(2)
				return set
			},
			2,
			true,
		},
		{
			"Given a set {1,2}, then run Search(3), false is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				set.Add(1)
				set.Add(2)
				return set
			},
			3,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := tt.initSet()
			search := set.Search(tt.search)

			if tt.want != search {
				t.Fatalf("Search(%v) = %v, want match for %v", tt.search, search, tt.want)
			}
		})
	}
}

func TestArrayBasedSetSize(t *testing.T) {
	tests := []struct {
		name    string
		initSet func() Set[int]
		want    int
	}{
		{
			"Given an empty set, then run Size(), 0 is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				return set
			},
			0,
		},
		{
			"Given a non-empty set with length 3, then run Size(), 3 is expected",
			func() Set[int] {
				set := NewArrayBasedSet[int]()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				return set
			},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := tt.initSet()
			size := set.Size()

			if tt.want != size {
				t.Fatalf("Size() = %v, want match for %v", size, tt.want)
			}
		})
	}

}

func TestArrayBasedSetUnion(t *testing.T) {
	tests := []struct {
		name    string
		initSets func() (Set[int], Set[int])
		want    []int
	}{
		{
			"{} + {} = {}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set2 := NewArrayBasedSet[int]()
				return set1, set2
			},
			nil,
		},
		{
			"{} + {1,2} = {1,2}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				set2.Add(2)
				return set1, set2
			},
			[]int{1,2},
		},
		{
			"{1,2,3} + {1,2} = {1,2,3}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(1)
				set1.Add(2)
				set1.Add(3)
				set2 := NewArrayBasedSet[int]()
				set2.Add(1)
				set2.Add(2)
				return set1, set2
			},
			[]int{1,2,3},
		},
		{
			"{1,2,3} + {4,5,6} = {1,2,3,4,5,6}",
			func() (Set[int], Set[int]) {
				set1 := NewArrayBasedSet[int]()
				set1.Add(1)
				set1.Add(2)
				set1.Add(3)
				set2 := NewArrayBasedSet[int]()
				set2.Add(4)
				set2.Add(5)
				set2.Add(6)
				return set1, set2
			},
			[]int{1,2,3,4,5,6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set1, set2 := tt.initSets()
			set := set1.Union(set2)

			members := set.Members()
			slices.Sort(members)
			if !reflect.DeepEqual(tt.want, members) {
				t.Fatalf("Union() = %v, want match for %v", members, tt.want)
			}
		})
	}

}
