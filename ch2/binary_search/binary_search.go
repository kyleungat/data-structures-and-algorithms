package binary_search

import (
	"cmp"
)

// parameters: a sorted array in asc order, target value
// return: index(int), exist(bool)
func binarySearch[T cmp.Ordered](sortedArray []T, target T) (int, bool) {
	// First, establish the lower and upper bounds
	lower_bound := 0
	upper_bound := len(sortedArray) - 1

	if upper_bound == -1 {
		return -1, false 
	}
	
	// keep checking the middlemost value between the upper and lower bounds
	for lower_bound <= upper_bound {
		// get the midpoint
		mid_point := (lower_bound + upper_bound) / 2 

		// check the value
		mid_point_value := sortedArray[mid_point]

		// If the value at the midpoint equals to the target, return it
		if mid_point_value == target {
			return mid_point, true
		}

		// If not, we change the lower or upper bound based on whether we need
		// to guess higher or lower:
		if mid_point_value < target {
			lower_bound = mid_point + 1
		} else {
			upper_bound = mid_point - 1
		}
	}

	// If we've narrowed the bounds until they've reached each other, all elements are checked
	// nothing is found
	return -1, false
}