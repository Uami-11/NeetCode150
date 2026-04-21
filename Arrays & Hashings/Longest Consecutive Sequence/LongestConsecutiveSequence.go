package main

func longestConsecutive(nums []int) int {
	hashNums := make(map[int]bool)
	for _, num := range nums {
		hashNums[num] = true
	}

	starters := make(map[int][]int)
	for value := range hashNums {
		if !hashNums[value-1] {
			starters[value] = append(starters[value], value)
		}
	}

	for value := range starters {
		next := value + 1
		for hashNums[next] {
			starters[value] = append(starters[value], next)
			next++
		}
	}

	var max int
	for _, sequence := range starters {
		if len(sequence) > max {
			max = len(sequence)
		}
	}

	return max
}
