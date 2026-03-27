package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		difference := target - num
		if _, exists := numMap[difference]; exists {
			return []int{numMap[difference], i}
		}
		numMap[num] = i
	}

	return []int{}
}
