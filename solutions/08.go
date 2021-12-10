package solutions

import (
	"fmt"
	"strconv"
)

// Task08 - 8.	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func Task08() {
	// set number and position
	var number int64 = 8
	var i uint = 0

	// get result
	resOne := setBitToOne(number, i)
	resZero := setBitToZero(number, i)
	fmt.Printf("%d - default number \nchanging bit: %d\n%d - change bit to one\n%d - change bit to zero\n\n", number, i, resOne, resZero)

	// get numbers in binary format

	//s := strconv.FormatInt(n, 2)
	//fmt.Printf("%064s", s)
	numberBinary := strconv.FormatInt(number, 2)
	resOneBinary := strconv.FormatInt(int64(resOne), 2)
	resZeroBinary := strconv.FormatInt(int64(resZero), 2)

	fmt.Printf("Result in tinary format:\n%s - default number \n%v - change bit to one\n%v - change bit to zero\n", numberBinary, resOneBinary, resZeroBinary)

	//fmt.Printf("\ndefault: %s\nsetBit: %s\nclearBit: %s\n",
	//	strconv.FormatInt(int64(number), 2),
	//	strconv.FormatInt(int64(resOne), 2),
	//	strconv.FormatInt(int64(resZero), 2))
}

func setBitToOne(number int64, position uint) int64 {
	// 1. "1 << position" - shift the number 1 the specified number of spaces in the integer (it becomes 0010, 0100, etc).
	// 2. "|" - OR with original input. This leaves the other bits unaffected but will always set the target bit to 1.
	number |= 1 << position
	return number
}

func setBitToZero(number int64, position uint) int64 {
	// 1. "1 << position" - shift the number 1 the specified number of spaces in the integer (it becomes 0010, 0100, etc).
	// 2. "^" - flip every bit in the mask with the ^ operator (0010 ->  1101).
	// 3. "&" - bitwise AND, which doesn't touch the numbers AND'ed with 1, but which will unset the value in the mask which is set to 0.
	mask := ^(1 << position)
	number &= int64(mask)
	return number
}
