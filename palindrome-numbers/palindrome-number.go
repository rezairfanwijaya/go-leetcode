package palindromenumbers

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	// x to sting
	xString := strconv.Itoa(x)

	// palindrome
	var palindrome string
	for i := len(xString) - 1; i >= 0; i-- {
		palindrome += string(xString[i])
	}

	// cek palindrome
	return xString == palindrome
}
