package main

import (
	"fmt"
	"os"
)

const USD_EUR = 0.8626
const USD_RUB = 81.91
const EUR_RUB = 97.29
const EUR_USD = 1.16
const RUB_EUR = 0.010279
const RUB_USD = 0.011905

func main() {
	for {
		var convertAgain string
		fmt.Println("Добро пожаловать в наш конвертор валют !")
		fmt.Println("Поддерживаемые валюты USD, EUR, RUB")
		calculateFrom, calculateIn, calculateHowMuch := getUserInput()
		convertResult := convertCurrency(calculateFrom, calculateIn, calculateHowMuch)
		if err := validateCurrency(calculateFrom); err != nil {
			fmt.Println("Ошибка", err)
		}
		if err := validateCurrency(calculateIn); err != nil {
			fmt.Println("Ошибка", err)
		}
		if err := validateNumbers(calculateHowMuch); err != nil {
			fmt.Println("Ошибка", err)

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
	for {
		fmt.Println("Введите исходную валюту(USD,EUR,RUB)")
		fmt.Scan(&calculateFrom)
		if calculateFrom == "stop" || calculateFrom == "exit" {
			fmt.Println("Выход по запросу")
			os.Exit(0)
		}
		if err := validateCurrency(calculateFrom); err == nil {
			break
		}
		fmt.Println("Ошибка, попробуйте снова")
	}
	for {
		fmt.Println("Введите в какую валюту вы желаете конвертировать(USD,EUR,RUB)")
		fmt.Scan(&calculateIn)
		if calculateIn == "stop" || calculateIn == "exit" {
			fmt.Println("Выход по запросу")
			os.Exit(0)
		}
		if calculateFrom == calculateIn {
			continue
		}
		if err := validateCurrency(calculateIn); err == nil {
			break
		}
		fmt.Println("Ошибка, попробуйте снова")
	}
	for {
		fmt.Println("Введите количество конвертируемой валюты")
		fmt.Scan(&calculateHowMuch)
		if err := validateNumbers(calculateHowMuch); err == nil {
			break
		}
		fmt.Println("Ошибка, попробуйте снова")
	}
	return
}
func validateNumbers(calculateHowMuch int) error {
	if calculateHowMuch <= 0 {
		return fmt.Errorf("Ошибка, число должно быть больше нуля")
	}
	return nil
}
func validateCurrency(currency string) error {
	if currency == "" {
		return fmt.Errorf("валюта не может быть пустой")
	}
	allowed := map[string]bool{"USD": true, "EUR": true, "RUB": true}
	if !allowed[currency] {
		return fmt.Errorf("допустимые валюты: USD, EUR, RUB")
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
	} else if firstCurrency == "RUB" && secondCurrency == "EUR" {
		result = float64(Quanity) * RUB_EUR
	} else if firstCurrency == "RUB" && secondCurrency == "USD" {
		result = float64(Quanity) * RUB_USD
	}
	return result
}
