package versionsort

func Compare(a, b string) int {
	return verrevcmp([]byte(a), []byte(b))
}

func verrevcmp(s1, s2 []byte) int {
	s1Len := len(s1)
	s2Len := len(s2)
	s1Pos := 0
	s2Pos := 0

	for s1Pos < s1Len || s2Pos < s2Len {
		first_diff := 0

		for (s1Pos < s1Len && !isDigit(s1[s1Pos])) || (s2Pos < s2Len && !isDigit(s2[s2Pos])) {
			var s1c byte
			var s2c byte

			if s1Pos < s1Len {
				s1c = s1[s1Pos]
			}
			if s2Pos < s2Len {
				s2c = s2[s2Pos]
			}

			s1Order := order(s1c)
			s2Order := order(s2c)

			if s1Order != s2Order {
				return s1Order - s2Order
			}

			s1Pos++
			s2Pos++
		}

		for s1Pos < s1Len && s1[s1Pos] == '0' {
			s1Pos++
		}
		for s2Pos < s2Len && s2[s2Pos] == '0' {
			s2Pos++
		}

		for s1Pos < s1Len && s2Pos < s2Len && isDigit(s1[s1Pos]) && isDigit(s2[s2Pos]) {
			if first_diff == 0 {
				first_diff = int(s1[s1Pos]) - int(s2[s2Pos])
			}
			s1Pos++
			s2Pos++
		}

		if s1Pos < s1Len && isDigit(s1[s1Pos]) {
			return 1
		}
		if s2Pos < s2Len && isDigit(s2[s2Pos]) {
			return -1
		}
		if first_diff != 0 {
			return first_diff
		}
	}

	return 0
}

func order(c byte) int {
	if isDigit(c) {
		return 0
	} else if isAlpha(c) {
		return int(c)
	} else if c == '~' {
		return -1
	} else if c == 0 {
		return 0
	} else {
		return int(c) + 256
	}
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func Sort(items []string) {
	if len(items) <= 1 {
		return
	}
	quickSort(items, 0, len(items)-1)
}

func quickSort(items []string, low, high int) {
	if low < high {
		p := partition(items, low, high)
		quickSort(items, low, p-1)
		quickSort(items, p+1, high)
	}
}

func partition(items []string, low, high int) int {
	pivot := items[high]
	i := low - 1
	for j := low; j < high; j++ {
		if Compare(items[j], pivot) < 0 {
			i++
			items[i], items[j] = items[j], items[i]
		}
	}
	items[i+1], items[high] = items[high], items[i+1]
	return i + 1
}
