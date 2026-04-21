package main

func maximumPossibleSize(nums []int) int {
	top := -1
	length := 0

	for _, num := range nums {
		if top <= num {
			length++
			top = num
		}
	}

	return length
}
