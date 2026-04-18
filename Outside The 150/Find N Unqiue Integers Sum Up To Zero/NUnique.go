package main

func sumZero(n int) []int {
	var arr []int
	var max int
	counter := 1

	if n%2 == 0 {
		max = n
	} else {
		max = n - 1
	}

	for i := range max {
		if i == max/2 {
			counter = -1
		}
		arr = append(arr, counter)

		if counter > 0 {
			counter++
		} else {
			counter--
		}
	}

	if max != n {
		arr = append(arr, 0)
	}

	return arr
}
