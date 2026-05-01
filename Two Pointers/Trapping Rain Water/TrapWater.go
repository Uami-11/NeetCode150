package main

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft := height[left]
	maxRight := height[right]

	var amount int
	var water int
	for left < right {
		if maxLeft <= maxRight {
			left++
			amount = maxLeft - height[left]
			amount = max(0, amount)
			water += amount
			if height[left] > maxLeft {
				maxLeft = height[left]
			}
		} else {
			right--
			amount = maxRight - height[right]
			amount = max(0, amount)
			water += amount
			if height[right] > maxRight {
				maxRight = height[right]
			}
		}
	}
	return water
}
