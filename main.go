package main

import (
	"fmt"
)

const (
	usdToEur = 0.87
	usdToRub = 78.74
	//  расчетах не участвует
	_eurToRub = usdToRub / usdToEur
)

const (
	usdCurrencyName = "USD"
	eurCurrencyName = "EUR"
	rubCurrencyName = "RUB"
	currencyNames   = usdCurrencyName + " " + eurCurrencyName + " " + rubCurrencyName
)

func main() {
	sourceCurrencyName := getSourceCurrencyName()
	valueToCalculate := getValueToCalculate()
	resultCurrencyName := getResultCurrencyName(sourceCurrencyName)

	result := calculate(valueToCalculate, sourceCurrencyName, resultCurrencyName)

	fmt.Println("_____________________________")
	fmt.Printf("Результат перевода %f %s в %s равен %.2f", valueToCalculate, sourceCurrencyName, resultCurrencyName, result)
}

func calculate(value float64, sourceCurrencyName string, resultCurrencyName string) float64 {
	divisor := 1.0

	if sourceCurrencyName == eurCurrencyName {
		divisor = usdToEur
	} else if sourceCurrencyName == rubCurrencyName {
		divisor = usdToRub
	}

	valueInUsd := value / divisor

	switch resultCurrencyName {
	case eurCurrencyName:
		return valueInUsd * usdToEur
	case rubCurrencyName:
		return valueInUsd * usdToRub
	default:
		return valueInUsd
	}
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
	switch unavailableCurrencyName {
	case usdCurrencyName:
		return eurCurrencyName, rubCurrencyName
	case eurCurrencyName:
		return usdCurrencyName, rubCurrencyName
	default:
		return usdCurrencyName, eurCurrencyName
	}
}
