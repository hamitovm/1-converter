package main

import (
	"fmt"
)

const (
	usdCurrencyName = "USD"
	eurCurrencyName = "EUR"
	rubCurrencyName = "RUB"
	currencyNames   = usdCurrencyName + " " + eurCurrencyName + " " + rubCurrencyName
)

var exchangeRateMap = map[string]float64{
	usdCurrencyName: 1,
	eurCurrencyName: 0.87,
	rubCurrencyName: 78.74,
}

func main() {
	sourceCurrencyName := getSourceCurrencyName()
	valueToCalculate := getValueToCalculate()
	resultCurrencyName := getResultCurrencyName(sourceCurrencyName)

	result := calculate(valueToCalculate, sourceCurrencyName, resultCurrencyName)

	fmt.Println("_____________________________")
	fmt.Printf("Результат перевода %f %s в %s равен %.2f", valueToCalculate, sourceCurrencyName, resultCurrencyName, result)
}

func calculate(value float64, sourceCurrencyName string, resultCurrencyName string) float64 {
	exchangeRate := exchangeRateMap[sourceCurrencyName]

	valueInUsd := value / exchangeRate

	return valueInUsd * exchangeRateMap[resultCurrencyName]
}

func requestReEntry() {
	fmt.Println("Введенное значение некорректно, введите еще раз:")
}

func getSourceCurrencyName() string {
	var sourceCurrencyName string

	fmt.Println("Введите название исходной валюты (", currencyNames, ") :")

	for {
		_, err := fmt.Scan(&sourceCurrencyName)

		if err != nil || (sourceCurrencyName != usdCurrencyName && sourceCurrencyName != eurCurrencyName && sourceCurrencyName != rubCurrencyName) {
			requestReEntry()
			continue
		}
		break
	}

	return sourceCurrencyName
}

func getValueToCalculate() float64 {
	var value float64

	fmt.Println("Введите сумму для конвертации:")

	for {
		_, err := fmt.Scan(&value)

		if err != nil || value < 0 {
			requestReEntry()
			continue
		}
		break
	}
	return value
}

func getResultCurrencyName(sourceCurrencyName string) string {
	aCurrencyName, bCurrencyName := getAvailableCurrencyNames(sourceCurrencyName)

	var resultCurrencyName string

	fmt.Println("Введите название целевой валюты для конвертации (", aCurrencyName, bCurrencyName, ") :")

	for {
		_, err := fmt.Scan(&resultCurrencyName)

		if err != nil || (resultCurrencyName != aCurrencyName && resultCurrencyName != bCurrencyName) {
			requestReEntry()
			continue
		}
		break
	}
	return resultCurrencyName
}

func getAvailableCurrencyNames(unavailableCurrencyName string) (string, string) {
	availableCurrencyNames := []string{}
	for key, _ := range exchangeRateMap {
		if key != unavailableCurrencyName {
			availableCurrencyNames = append(availableCurrencyNames, key)
		}
	}
	return availableCurrencyNames[0], availableCurrencyNames[1]
}
