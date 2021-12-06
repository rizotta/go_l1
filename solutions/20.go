package solutions

import (
	"fmt"
	"strings"
)

// Task20 - 20.	Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
func Task20() {
	str := "snow dog sun"
	fmt.Println(str)
	fmt.Println(reverseWords(str))
}

func reverseWords(words string) string {
	// split string to array
	arr := strings.Split(words, " ") // [snow dog sun]

	arrLen := len(arr) // ex. 3

	// swap words in array by index
	for i := 0; i < arrLen/2; i++ {
		arr[i], arr[arrLen-1-i] = arr[arrLen-1-i], arr[i] // ex. arr[0], arr[3-1-0] = arr[2], arr[0]
	}

	// convert array elements to string
	return strings.Join(arr, " ")
}
