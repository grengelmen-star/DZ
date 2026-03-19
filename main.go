package main

import "fmt"

func main() {
	const USD_EUR = 0.8667
	const USD_RUB = 81.91
	EUR_RUB := USD_RUB * USD_EUR
	fmt.Println(EUR_RUB)

}
func getUserInput() {
	var calculateFrom string
	var calculateIn string
	var calculateHowMuch int
	fmt.Println("Введите исходную валюту")
	fmt.Scan(&calculateFrom)
	fmt.Println("Введите в какую валюту вы желаете конвертировать")
	fmt.Scan(&calculateIn)
	fmt.Println("Введите количество конвертируемой валюты")
	fmt.Scan(&calculateHowMuch)

}
func convertCurrency(firstCurrency string, secondCurrency string, Quanity int) {}
