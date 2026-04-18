// This does not work because slices.Delete is also O(n)
package main

import "slices"

func maximumPossibleSize(nums []int) int {
	for i := 0; i < len(nums)-1; {
		if nums[i] > nums[i+1] {
			nums = slices.Delete(nums, i+1, i+2)
		} else {
			i++
		}
	}

	return len(nums)
}
