package main

import (
	"fmt"
	"l1/solutions"
	"strconv"
)

func main() {
	//taskNum, _ := getTask()
	taskNum := 12
	switch taskNum {
	default:
		fmt.Println("wrong number")
		return
	case 1:
		solutions.Task01()
	case 2:
		solutions.Task02()
	case 3:
		solutions.Task03()
	case 4:
		solutions.Task04()
	case 5:
		solutions.Task05()
	case 6:
		solutions.Task06()
	case 7:
		solutions.Task07()
	case 8:
		solutions.Task08()
	case 9:
		solutions.Task09()
	case 10:
		solutions.Task10()
	case 11:
		solutions.Task11()
	case 12:
		solutions.Task12()
	case 13:
		solutions.Task13()
	case 14:
		solutions.Task14() // empty
	case 15:
		solutions.Task15() // empty
	case 16:
		solutions.Task16()
	case 17:
		solutions.Task17()
	case 18:
		solutions.Task18() // empty
	case 19:
		solutions.Task19()
	case 20:
		solutions.Task20()
	case 21:
		solutions.Task21() // empty
	case 22:
		solutions.Task22() // empty
	case 23:
		solutions.Task23()
	case 24:
		solutions.Task24() // empty
	case 25:
		solutions.Task25()
	case 26:
		solutions.Task26()
	}
}

// getTask - request task number in console
func getTask() (int, error) {
	var taskStr string
	fmt.Print("Enter number of task from 1 to 26 => ")
	_, err := fmt.Scan(&taskStr)
	if err != nil {
		return 0, err
	}
	taskNum, _ := strconv.Atoi(taskStr)
	return taskNum, nil
}
