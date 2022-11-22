package mergetwosortedlists

func MergeTwoSortedList(a []int, b []int) []int {
	a = append(a, b...)
	return Sorting(a)
}

func Sorting(a []int) []int {
	for i := 0; i < len(a); i++ {
		tmp := a[i]
		j := i
		for j > 0 && a[j-1] > tmp {
			a[j] = a[j-1]
			j--
		}
		a[j] = tmp

	}
	return a
}
