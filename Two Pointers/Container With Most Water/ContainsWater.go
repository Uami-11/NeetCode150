package main

func maxArea(heights []int) int {
	var water int
	left := 0
	right := len(heights) - 1
	for left < right {
		currentContainer := min(heights[left], heights[right]) * (right - left)
		if currentContainer > water {
			water = currentContainer
		}

		if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}
	return water
}
