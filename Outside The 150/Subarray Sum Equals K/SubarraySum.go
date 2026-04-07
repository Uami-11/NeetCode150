package main

func subarraySum(nums []int, k int) int {
	seenSums := make(map[int]int)
	var prefixSum int
	var count int

	seenSums[0] = 1

	for _, num := range nums {
		prefixSum += num
		if value, exists := seenSums[prefixSum-k]; exists {
			count += value
		}
		seenSums[prefixSum]++
	}

	return count
}
