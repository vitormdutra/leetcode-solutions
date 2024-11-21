package main

import (
	"fmt"
	"slices"
)

// You are given two arrays of equal length, nums1 and nums2.

// Each element in nums1 has been increased (or decreased in the case of negative) by an integer, represented by the variable x.

// As a result, nums1 becomes equal to nums2. Two arrays are considered equal when they contain the same integers with the same frequencies.

// Return the integer x.

func addedInteger(nums1 []int, nums2 []int) int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	result := nums2[0] - nums1[0]
	fmt.Println(result)

	return result
}

func main() {
	addedInteger([]int{2, 6, 4}, []int{9, 7, 5})
	addedInteger([]int{10}, []int{5})
	addedInteger([]int{1, 1, 1, 1}, []int{1, 1, 1, 1})
}
