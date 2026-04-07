package main

func productExceptSelf(nums []int) []int {
	var result []int
	for range nums {
		result = append(result, 1)
	}

	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
