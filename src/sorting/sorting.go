/*
Author: John Lahut
Project: Sorting
Date: 9/18/2018
*/
package sorting

import "time"

// MergeSort docs
func MergeSort(arr []int) (t int64) {
	start := time.Now()

	// base case - if we have an slice of length 1
	if len(arr) > 1 {

		// compute midpoint, and create new slices (need to copy because s slice of a slice uses same memory)
		mid := len(arr) / 2
		left := make([]int, len(arr[:mid]))
		right := make([]int, len(arr[mid:]))
		copy(left, arr[:mid])
		copy(right, arr[mid:])

		// sort our smaller slices
		MergeSort(left)
		MergeSort(right)

		// i, j trace the sub slices
		// k traces the master slice
		i, j, k := 0, 0, 0

		// loop through placing lists in order
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				arr[k] = left[i]
				i++
			} else {
				arr[k] = right[j]
				j++
			}
			k++
		}

		// append any left over elements
		for ; i < len(left); i++ {
			arr[k] = left[i]
			k++
		}
		for ; j < len(right); j++ {
			arr[k] = right[j]
			k++
		}
	}
	// return time elapsed
	return int64(time.Since(start) / time.Millisecond)
}
