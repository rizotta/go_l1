package solutions

import (
	"fmt"
)

// Task12 - 12.	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func Task12() {
	strArr := []string{"cat", "cat", "dog", "cat", "tree"}
	strMap := make(map[string]bool)
	var res []string
	for _, v := range strArr {
		strMap[v] = true
	}
	for k := range strMap {
		res = append(res, k)
	}

	fmt.Println(res)
}
