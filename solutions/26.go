package solutions

import (
	"fmt"
	"strings"
)

// Task26 - 26.	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// Например: abcd — true, abCdefAaf — false, aabcd — false
func Task26() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "abcdA"

	fmt.Printf("%s - %t\n", str1, isUnique(str1))
	fmt.Printf("%s - %t\n", str2, isUnique(str2))
	fmt.Printf("%s - %t\n", str3, isUnique(str3))
}

func isUnique(str string) bool {
	str = strings.ToLower(str)
	lenStr := len(str)

	for i := 0; i < lenStr; i++ {
		for j := i + 1; j < lenStr; j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}
