package main

import (
	"fmt"
	"strconv"
)

func main() {
	var plusOneInput = []int{1, 2, 3}
	fmt.Println("Result of plusOne: ", plusOne(plusOneInput))
}

func plusOne(digits []int) []int {

	strDigit := ""
	for i := 0; i < len(digits); i++ {
		strDigit = strDigit + strconv.Itoa(digits[i])
	}

	intDigit, _ := strconv.Atoi(strDigit)
	intDigit = intDigit + 1

	strTotalSum := strconv.Itoa(intDigit)

	result := make([]int, len(strTotalSum))
	for i := 0; i < len(strTotalSum); i++ {
		intVar, _ := strconv.Atoi(string(strTotalSum[i]))
		result[i] = intVar
	}
	return result
}
