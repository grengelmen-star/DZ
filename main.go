package main

import (
	"fmt"
	"strconv"
)

const USD_EUR = 0.8626
const USD_RUB = 81.91
const EUR_RUB = 97.29
const EUR_USD = 1.16

func main() {
	for {
		var convertAgain string
		calculateFrom, calculateIn, calculateHowMuch := getUserInput()
		convertResult := convertCurrency(calculateFrom, calculateIn, calculateHowMuch)
		if err := validateCurrencies(calculateFrom, calculateIn); err != nil {
			fmt.Println("Ошибка", err)
			return
		}
		fmt.Printf("Вы конвертировали: %d %s в %s результат : %.2f \n", calculateHowMuch, calculateFrom, calculateIn, convertResult)
		fmt.Println("Желаете конвертировать еще раз ?")
		fmt.Scan(&convertAgain)
		if calculateFrom == "stop" || calculateIn == "stop" {
			fmt.Println("Спасибо за использование нашей программы")
			break
		} else if convertAgain == "N" || convertAgain == "No" {
			fmt.Println("Спасибо за использование нашей программы")
			break
		} else {
			fmt.Println("Конвертируем вновь !")
		}
	}
}
func getUserInput() (calculateFrom string, calculateIn string, calculateHowMuch int) {
	fmt.Println("Введите исходную валюту")
	fmt.Scan(&calculateFrom)
	fmt.Println("Введите в какую валюту вы желаете конвертировать")
	fmt.Scan(&calculateIn)
	fmt.Println("Введите количество конвертируемой валюты")
	fmt.Scan(&calculateHowMuch)
	return
}
func validateCurrencies(calculateFrom string, calculateIn string) error {
	if calculateFrom == "" || calculateIn == "" {
		return fmt.Errorf("Валюты не могут быть пустыми")
	}
	if calculateFrom == calculateIn {
		return fmt.Errorf("Вы не можете конвертировать %s в %s", calculateFrom, calculateIn)
	}
	if _, err := strconv.Atoi(calculateFrom); err == nil {
		return fmt.Errorf("Ошибка, нужно ввести валюту")
	}
	if _, err := strconv.Atoi(calculateIn); err == nil {
		return fmt.Errorf("Ошибка, нужно ввести валюту")
	}
	return nil
}
func convertCurrency(firstCurrency string, secondCurrency string, Quanity int) float64 {
	var result float64
	if firstCurrency == "USD" && secondCurrency == "EUR" {
		result = float64(Quanity) * USD_EUR
	} else if firstCurrency == "USD" && secondCurrency == "RUB" {
		result = float64(Quanity) * USD_RUB
	} else if firstCurrency == "EUR" && secondCurrency == "RUB" {
		result = float64(Quanity) * EUR_RUB
	} else if firstCurrency == "EUR" && secondCurrency == "USD" {
		result = float64(Quanity) * EUR_USD
	}
	return result
}
