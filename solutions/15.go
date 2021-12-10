package solutions

import "fmt"

/*
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/

// Task15 - 15.	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации
func Task15() {
	/*
		   	1. It's not a good idea to use global variable ("justString") when it is not necessary.
		   	Possible conflicts in the future, confusion, the code is less readable.
		   	It is necessary to use the variable in the "main()" function.

			2. "createHugeString" function is not implemented.
		   	The program won't work.
		   	The function needs to be implemented.

		   	3. Unused variable 'justString'.
		   	Build error
			Use it or delete variable.
	*/

	// Example of working code
	var justString string
	someFunc(&justString)
	fmt.Println(justString)
}

func someFunc(justString *string) {
	v := createHugeString(1 << 10)
	*justString = v[:100]
}

func createHugeString(size int) string {
	var res string
	for i := 0; i < size; i++ {
		res += "Text"
	}
	return res
}
