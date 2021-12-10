package solutions

import (
	"fmt"
)

// Task14 - 14.	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
func Task14() {
	ch := make(chan int)
	arr := []interface{}{123, "some text", true, ch}

	for _, val := range arr {
		switch val.(type) {
		case int:
			fmt.Printf("int: %d\n", val)
		case string:
			fmt.Printf("string: %s\n", val)
		case bool:
			fmt.Printf("bool: %t\n", val)
		case chan int:
			fmt.Printf("channel: %T\n", val)
		}
	}
}
