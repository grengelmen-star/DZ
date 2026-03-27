package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var inputOperation string
var inputNumbers string

func main() {
	fmt.Println("-----Калькулятор-----")
	for {
		getUserInput()
		parts := strings.Split(inputNumbers, ",")
		var numbers []int
		for _, value := range parts {
			value = strings.TrimSpace(value)
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Ошибка преобразования `%s` : %v\n", value, err)
				continue
			}
			numbers = append(numbers, num)
		}
		valid := true
		switch inputOperation {
		case "AVG":
			fmt.Println("Результат подсчета AVG : ", calculateAVG(numbers))
		case "SUM":
			fmt.Println("Результат подсчета SUM : ", calculateSUM(numbers))
		case "MED":
			fmt.Println("Результат подсчета MED : ", calculateMED(numbers))
		default:
			fmt.Println("Ошибка: неизвестная операция, пожалуйста, повторите ввод")
			valid = false
		}
		if !valid {
			continue
		}
		fmt.Println("Желаете провести еще расчет ?(Y/N)")
		var continueOperation string
		fmt.Scan(&continueOperation)
		if continueOperation == "Y" || continueOperation == "Yes" || continueOperation == "yes" {
			continue
		} else {
			break
		}
	}

}

func getUserInput() {
	fmt.Println("Введите пожалуйста желаемую операцию(AVG - среднее арифметическое,SUM - сумма, MED - медиана):")
	fmt.Scan(&inputOperation)
	fmt.Println("Введите пожалуйста цифры для операции через запятую одной строкой:")
	fmt.Scan(&inputNumbers)
}
func calculateAVG(s []int) float64 {
	if len(s) == 0 {
		fmt.Println("Ошибка!")
		return 0
	}
	var cycleResult int
	for _, value := range s {
		cycleResult += value
	}
	result := float64(cycleResult) / float64(len(s))
	return result
}
func calculateSUM(s []int) int {
	if len(s) == 0 {
		fmt.Println("Ошибка!")
		return 0
	}
	var cycleResult int
	for _, value := range s {
		cycleResult += value
	}
	return cycleResult
}
func calculateMED(s []int) float64 {
	if len(s) == 0 {
		fmt.Println("Ошибка!")
		return 0
	}
	sorted := make([]int, len(s))
	copy(sorted, s)
	sort.Ints(sorted)
	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	}
	mid := n / 2
	return float64(sorted[mid-1]+sorted[mid]) / 2
}
