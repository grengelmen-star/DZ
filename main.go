package main

import "fmt"

func main() {
	const USD_EUR = 0.8667
	const USD_RUB = 81.91
	EUR_RUB := USD_RUB * USD_EUR
	fmt.Println(EUR_RUB)

}
