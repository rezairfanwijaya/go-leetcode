package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	b := false
	var str string
	for {
		var code string
		var waitCompare string
		for k, word := range strs {
			if k == 0 {
				if len(word) <= i {
					b = true
					break
				}
				waitCompare = string(word[i])
			} else {
				if len(word) <= i {
					b = true
					break
				}
				code = string(word[i])
				if code != waitCompare {
					b = true
					break
				}
			}
		}
		if b {
			break
		}
		waitCompare = code
		str += waitCompare
		i++
	}
	return str
}
