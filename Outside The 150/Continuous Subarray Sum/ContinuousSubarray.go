package main

func checkSubarraySum(nums []int, k int) bool {
	check := make(map[int]int)

	var prefix int

	check[0] = -1
	for i, num := range nums {
		prefix += num

		if _, exists := check[prefix%k]; exists {
			if i-check[prefix%k] >= 2 {
				return true
			}
		} else {
			check[prefix%k] = i
		}
	}

	return false
}
