package main

import (
	"fmt"
	"os"
)

var rates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.8626,
	"RUB": 81.91,
}

func main() {
	for {
		var convertAgain string
		fmt.Println("Добро пожаловать в наш конвертор валют !")
		fmt.Println("Поддерживаемые валюты USD, EUR, RUB")
		calculateFrom, calculateIn, calculateHowMuch := getUserInput()
		convertResult := convertCurrency(calculateFrom, calculateIn, calculateHowMuch, &rates)
		if err := validateCurrency(calculateFrom, &rates); err != nil {
			fmt.Println("Ошибка", err)
		}
		if err := validateCurrency(calculateIn, &rates); err != nil {
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
		if err := validateCurrency(calculateFrom, &rates); err == nil {
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
		if err := validateCurrency(calculateIn, &rates); err == nil {
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
func validateCurrency(currency string, ratesPoint *map[string]float64) error {
	if currency == "" {
		return fmt.Errorf("валюта не может быть пустой")
	}
	if _, ok := (*ratesPoint)[currency]; !ok {
		return fmt.Errorf("Ошибка, допустимые валюты USD,EUR,RUB")
	}
	return nil
}
func convertCurrency(firstCurrency string, secondCurrency string, Quanity int, ratesPoint *map[string]float64) float64 {
	usdValidate := float64(Quanity) / (*ratesPoint)[firstCurrency]
	return usdValidate * (*ratesPoint)[secondCurrency]
}
