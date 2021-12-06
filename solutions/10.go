package solutions

import "fmt"

// Task10 - 10.	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
func Task10() {
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// make groups
	groups := make(map[int][]float32)

	// fill groups
	for _, val := range temp {
		groupName := int(val/10) * 10                      // get groupName from temp, ex. 24.5 => 20
		groups[groupName] = append(groups[groupName], val) // add value in groups[groupName]
	}

	// show groups
	fmt.Println(groups)
}
