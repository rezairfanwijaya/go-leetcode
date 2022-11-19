package validparentheses

func IsValid(s string) bool {
	lengthS := len(s)

	// * akan false ketika length 0, 1 atau ganjil
	if lengthS == 0 || lengthS == 1 || lengthS%2 != 0 {
		return false
	}

	// * memisahkan open dan close
	parenthesesOpen := make(map[string]int)
	parenthesesClose := make(map[string]int)

	// * memasukan open dan close kedalam map
	for index, v := range s {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			parenthesesOpen[string(v)] = index
		} else {
			parenthesesClose[string(v)] = index
		}
	}

	// * cek open dan close dengan kondisi
	// * index open < index close
	for _, v := range s {
		if string(v) == "(" {
			isValid := checkIndex("(", ")", parenthesesOpen, parenthesesClose)
			if !isValid {
				return false
			}
		} else if string(v) == "{" {
			isValid := checkIndex("{", "}", parenthesesOpen, parenthesesClose)
			if !isValid {
				return false
			}
		} else if string(v) == "[" {
			isValid := checkIndex("{", "}", parenthesesOpen, parenthesesClose)
			if !isValid {
				return false
			}
		}
	}

	return true
}

// * index open harus lebih kecil dari index close
func checkIndex(
	openString, closeString string,
	openParentheses, closeParentheses map[string]int) bool {

	indexOpen := openParentheses[openString]
	indexClose := closeParentheses[closeString]
	return indexOpen < indexClose
}
