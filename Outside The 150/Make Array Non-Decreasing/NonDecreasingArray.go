package main

func maximumPossibleSize(nums []int) int {
	var stack []int

	for _, num := range nums {
		if len(stack) == 0 || stack[len(stack)-1] <= num {
			stack = append(stack, num)
		}
	}

	return len(stack)
}
