package addtwonumbers

func AddtwoNumber(nums []int, target int) []int {
	var result []int

	for index, v := range nums {
		for i := index + 1; i < len(nums); i++ {
			if v+nums[i] == target {
				result = append(result, index, i)
			}
		}
	}

	return result
}
