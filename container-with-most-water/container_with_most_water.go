package containerwithmostwater

func MaxAge(height []int) int {
	length := len(height)

	var currentIndex, result int

	lastIndex := length - 1

	for currentIndex < lastIndex {
		currentArea := min(height[currentIndex], height[lastIndex]) * (lastIndex - currentIndex)
		result = max(result, currentArea)

		if height[currentIndex] < height[lastIndex] {
			currentIndex++
		} else {
			lastIndex--
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
