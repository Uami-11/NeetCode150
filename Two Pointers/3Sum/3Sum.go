package main

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var triplets [][]int
	for i, num := range nums {
		if i != 0 && num == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum > -num {
				right--
			} else if sum < -num {
				left++
			} else {
				triplets = append(triplets, []int{num, nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return triplets
}
