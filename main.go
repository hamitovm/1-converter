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
