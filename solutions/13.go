package solutions

import "fmt"

// Task13 - 13.	Поменять местами два числа без создания временной переменной.
func Task13() {
	a := 2
	b := 3
	fmt.Printf("a: %d, b: %d - default\n", a, b)

	// case 1: simple change in golang with multiple assignment
	a, b = b, a
	fmt.Printf("a: %d, b: %d - first schange\n", a, b)

	// case 2: just math
	a = a + b // 2 + 3 = 5
	b = a - b // 5 - 2 = 3
	a = a - b // 5 - 3 = 2
	fmt.Printf("a: %d, b: %d - second schange\n", a, b)
}
