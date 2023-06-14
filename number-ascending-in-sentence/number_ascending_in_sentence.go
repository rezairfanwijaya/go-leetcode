package numberascendinginsentence

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	var numbersInString []int
	sArray := strings.Split(s, " ")

	for _, v := range sArray {
		numberString, err := strconv.Atoi(string(v))
		if err == nil {
			numbersInString = append(numbersInString, numberString)
		}
	}

	firstNumber := numbersInString[0]
	for i := 1; i < len(numbersInString); i++ {
		if firstNumber > numbersInString[i] || firstNumber == numbersInString[i] {
			return false
		}

		firstNumber = numbersInString[i]
	}

	return true
}
