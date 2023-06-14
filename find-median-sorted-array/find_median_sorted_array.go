package findmediansortedarray

func findMedianSortedArray(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	var median float64

	lengthNums1 := len(nums1)

	i := 1
	for i < len(nums1) {
		j := i
		for j >= 1 && nums1[j] < nums1[j-1] {
			nums1[j], nums1[j-1] = nums1[j-1], nums1[j]

			j--
		}

		i++
	}
	if len(nums1)%2 != 0 {
		indexMedian := lengthNums1 / 2
		median = float64(nums1[indexMedian])
	} else {
		firstNumber := lengthNums1 / 2
		median = (float64(nums1[firstNumber]) + float64(nums1[firstNumber-1])) / float64(2)
	}

	return median
}
