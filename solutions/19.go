package solutions

import (
	"fmt"
)

// Task19 - 19. Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.
func Task19() {
	str := "ГлавカРыбÄ"
	fmt.Println(str)
	fmt.Println(reverseString(str))
}

func reverseString(str string) string {
	// use rune(int32) to convert Unicode string
	runes := []rune(str) // ex. runes = [1043 1083 1072 1074 12459 1056 1099 1073 196]
	runLen := len(runes) // ex. runLen = 9

	// swap elements by index
	for i := 0; i < runLen/2; i++ {
		runes[i], runes[runLen-1-i] = runes[runLen-1-i], runes[i] // ex. runes[0], runes[9-1-0] = rune[8], rune[0]
	}
	return string(runes)
}
