package main

import "fmt"

const (
	usdToEur = 1.16
	usdToRub = 78.74
	eurToRub = usdToRub * usdToEur
)

func main() {
	fmt.Println(usdToEur, usdToRub, eurToRub)
}

func getUserInput() string {
	var inputData string
	fmt.Scan(&inputData)
	return inputData
}

func calculate(value float64, valueCurrency string, resultCurrency string) float64 {}
