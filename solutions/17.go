package solutions

import "fmt"

// Task17 - 17.	Реализовать бинарный поиск встроенными методами языка.
func Task17() {
	arr := []int{-7, -5, -4, -3, -1, 1, 3, 4, 5, 7}
	findNumber := -3
	result := binarySearch(arr, findNumber)
	if result < 0 {
		fmt.Printf("not found %d\n", findNumber)
	} else {
		fmt.Printf("arr[%d] = %d\n", result, findNumber)
	}
}

func binarySearch(arr []int, findNumber int) (result int) {
	// get middle of array
	mid := len(arr) / 2

	switch {
	case len(arr) == 0:
		// if len == 0 => can't find number
		result = -1
	case arr[mid] > findNumber:
		// if findNumber < middle => search to the left of the middle
		result = binarySearch(arr[:mid], findNumber)
	case arr[mid] < findNumber:
		// if findNumber > middle => search to the right of the middle
		result = binarySearch(arr[mid+1:], findNumber)
		// if the number is found, then add index of mic from other recursive function + 1 (because first index is 0, just in right side)
		if result >= 0 {
			result += mid + 1
		}
	default:
		result = mid // arr[mid] = findNumber, number is found
	}
	return result
}
