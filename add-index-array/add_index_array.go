package addindexarray

import (
	"fmt"
	"strconv"
)

const MAX = 10

func AddTwoNumbers(a, b []int) int {
	// panjang array harus sama
	if len(a) != len(b) {
		return 0
	}

	// array baru untuk hasil penjumlahan
	var resultTMP []int

	// untuk menampung nilai jika lebih dari 10
	tmpGratherThenTen := 0

	// jumlahkan index perama pada array pertama
	// dengan index pertama pada array kedua
	for i := 0; i < len(a); i++ {
		sum := a[i] + b[i] + tmpGratherThenTen
		if sum >= MAX {
			splitNumber := SplitNumber(sum)
			tmpGratherThenTen = splitNumber[0]
			resultTMP = append(resultTMP, splitNumber[1])
		} else {
			resultTMP = append(resultTMP, sum)
		}
	}

	// generate result
	return GenerateResult(resultTMP)
}

func SplitNumber(number int) []int {
	// ubah ke string
	numberOfString := strconv.Itoa(number)

	// ambil angka pertama dan kedua
	fisrtNumber, _ := strconv.Atoi(string(numberOfString[0]))
	SecondNumber, _ := strconv.Atoi(string(numberOfString[1]))

	return []int{fisrtNumber, SecondNumber}
}

func GenerateResult(slice []int) (result int) {
	tmpResult := ""

	for i := len(slice) - 1; i >= 0; i-- {
		tmp := fmt.Sprintf("%v", slice[i])
		tmpResult = tmpResult + tmp
	}

	result, _ = strconv.Atoi(tmpResult)

	return result
}
