package stringtointeger

import (
	"strconv"
	"strings"
)

func StringToInteger(s string) int {
	// result
	var myNumber int

	// split string
	stringSplit := strings.Split(s, " ")

	// is number in end string
	isInTheEndAfterWord := false

	mapIndex := make(map[string]int)

	// string checking
	for index, v := range stringSplit {
		// convert string to number
		num, err := strconv.Atoi(v)
		if err == nil { // if is number
			myNumber = num
			mapIndex["index number"] = index
		} else {
			if v != "" {
				mapIndex["index word"] = index
			} else {
				mapIndex["index word"] = 9999999999999
			}
		}

	}

	if mapIndex["index word"] < mapIndex["index number"] {
		isInTheEndAfterWord = true
	}

	if isInTheEndAfterWord {
		return 0
	}

	return myNumber
}
