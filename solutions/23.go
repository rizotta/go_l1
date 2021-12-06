package solutions

import "fmt"

// Task23 - 23.	Удалить i-ый элемент из слайса.
func Task23() {
	sl := []int{-7, -5, -4, -3, -1, 1, 3, 4, 5, 7}
	i := 3
	fmt.Println(sl, " - original slice")
	slNew := removeByIndex(sl, i)
	fmt.Printf("%v - slice with delete element sl[%d]\n", slNew, i)
}

func removeByIndex(s []int, i int) []int {
	return append(s[:i], s[i+1:]...) // [-7 -5 -4] and [-1 1 3 4 5 7]
}
