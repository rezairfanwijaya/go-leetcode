package validparentheses

func IsValid(s string) bool {
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	stack := []rune{}
	opener := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, letter := range s {
		if closer, ok := opener[letter]; ok {
			stack = append(stack, closer)
			continue
		}

		idx := len(stack) - 1
		if idx < 0 || stack[idx] != letter {
			return false
		}

		stack = stack[:idx]
	}

	return len(stack) == 0
}
