package solutions

import "fmt"

// Task16 - 16.	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
func Task16() {
	arr := []int{1, 7, 3, 5, -3, 4, -1, -7, -5, -4}
	fmt.Println(arr, " - default")
	quicksort(arr)
	fmt.Println(arr, " - sort")
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1 // ex. 0, 9

	// select pivot
	pivot := len(arr) / 2 // ex. 10/2 = 5, arr[pivot] = 4

	// move pivot to the right
	arr[pivot], arr[right] = arr[right], arr[pivot] // ex. arr[pivot] = arr[9] = -4, ex. arr[right] = arr[5] = 4

	// pile elements smaller than pivot on the left
	for k := range arr {
		if arr[k] < arr[right] {
			arr[left], arr[k] = arr[k], arr[left]
			left++
		}
	} // ex. res: [1 3 -3 -4 -1 -7 -5 5 7 4] , where arr[pivot] = 5

	// place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left] // ex. res: [1 3 -3 -4 -1 -7 -5 4 7 5]

	// arr[:left] - elements <= pivot
	quicksort(arr[:left]) // Step 1 get [1 3 -3 -4 -1 -7 -5], arr[pivot] = -5, etc.

	// arr[:left] - elements > pivot
	quicksort(arr[left+1:])
	return arr
}

// https://qvault.io/golang/quick-sort-golang/
// https://stackoverflow.com/questions/15802890/idiomatic-quicksort-in-go/15803401
