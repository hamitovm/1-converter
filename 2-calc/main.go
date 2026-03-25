package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var operationTypes = []string{"avg", "sum", "med"}

func main() {
	operationType := getOperationType()
	numbers := getNumbers()

	switch operationType {
	case "sum":
		fmt.Println("Сумма равна", getSumValue(numbers))
	case "avg":
		fmt.Println("Средняя равна", getAverageValue(numbers))
	case "med":
		fmt.Println("Медиана равна", getMedianValue(numbers))

	}
}

func readUserInput() (string, error) {
	input, err := reader.ReadString('\n')
	input = strings.TrimRight(input, "\r\n")
	return input, err
}

func getOperationType() string {

	operationTypesPrintValue := strings.Join(operationTypes, ", ")

	fmt.Println("Введите тип операции (", operationTypesPrintValue, "): ")
	operationType, _ := readUserInput()

	for {
		isValidValue := slices.Contains(operationTypes, operationType)
		if !isValidValue {
			fmt.Printf("Вы ввели некорректное значение, значение должно быть одним из следующих: %s. Попробуйте еще раз: ", operationTypesPrintValue)
			operationType, _ = readUserInput()
			continue
		}
		break
	}

	return operationType
}

func getNumbers() []int {
	fmt.Println("Введите значения:")
	input, _ := readUserInput()

	var numbers []int

	parts := strings.Split(input, ",")

	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		num, err := strconv.Atoi(trimmedPart)
		if err != nil {
			fmt.Println("Обнаружено нечисловое значение")
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func getAverageValue(numbers []int) float64 {
	sum := getSumValue(numbers)
	return float64(sum / len(numbers))
}

func getSumValue(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func getMedianValue(numbers []int) float64 {
	sort.Ints(numbers)
	length := len(numbers)
	if length%2 == 0 {
		middleIndex := length / 2
		firstValue := numbers[middleIndex]
		secondValue := numbers[middleIndex-1]
		return float64((firstValue + secondValue) / 2)
	}
	return float64(numbers[length/2])
}
