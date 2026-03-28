package main

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	frequency := [][]int{}
	for i := 0; i <= len(nums); i++ {
		var empty []int
		frequency = append(frequency, empty)
	}

	for num, freq := range count {
		frequency[freq] = append(frequency[freq], num)
	}

	var result []int
	for i := len(frequency) - 1; i >= 0; i-- {
		if len(frequency[i]) != 0 {
			for _, top := range frequency[i] {
				result = append(result, top)
				if len(result) == k {
					return result
				}
			}
		}
	}

	return result
}
