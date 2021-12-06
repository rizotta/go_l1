package solutions

import "fmt"

// Task11 - 11.	Реализовать пересечение двух неупорядоченных множеств.
func Task11() {
	arr1 := []int{9, -7, 5, -3, 1, -1, 3, -5, 7, -9, -11}
	arr2 := []int{10, -11, 7, -4, 1, -2, 5, -8}
	fmt.Println(intersect(arr1, arr2))
}

func intersect(arr1, arr2 []int) []int {
	var res []int // create intersection result
	for _, val1 := range arr1 {
		for _, val2 := range arr2 {
			if val1 == val2 { // if values are identical
				res = append(res, val2) // push value in intersection result
			}
		}
	}
	return res
}
