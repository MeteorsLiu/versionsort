package versionsort

// Compare compares two version strings and returns:
// -1 if a < b
//  0 if a == b
//  1 if a > b
func Compare(a, b string) int {
	return verrevcmp([]byte(a), []byte(b))
}

// verrevcmp implements version number comparison algorithm (similar to GNU strverscmp)
// Compares character and numeric segments separately, with numeric segments compared by value
func verrevcmp(s1, s2 []byte) int {
	s1Len := len(s1)
	s2Len := len(s2)
	s1Pos := 0
	s2Pos := 0

	for s1Pos < s1Len || s2Pos < s2Len {
		first_diff := 0

		// Compare non-digit portions
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

		// Skip leading zeros
		for s1Pos < s1Len && s1[s1Pos] == '0' {
			s1Pos++
		}
		for s2Pos < s2Len && s2[s2Pos] == '0' {
			s2Pos++
		}

		// Compare numeric portions
		for s1Pos < s1Len && s2Pos < s2Len && isDigit(s1[s1Pos]) && isDigit(s2[s2Pos]) {
			if first_diff == 0 {
				first_diff = int(s1[s1Pos]) - int(s2[s2Pos])
			}
			s1Pos++
			s2Pos++
		}

		// If one numeric segment is longer, that number is larger
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

// order returns the sorting priority of a character
// digits: 0, letters: ASCII value, '~': -1, null: 0, others: ASCII value + 256
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

// isDigit checks if the character is a digit
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// isAlpha checks if the character is a letter
func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// Sort sorts a slice of strings by version number (in-place sorting)
func Sort(items []string) {
	if len(items) <= 1 {
		return
	}
	quickSort(items, 0, len(items)-1)
}

// quickSort implements the quicksort algorithm
func quickSort(items []string, low, high int) {
	if low < high {
		p := partition(items, low, high)
		quickSort(items, low, p-1)
		quickSort(items, p+1, high)
	}
}

// partition is the partition function for quicksort
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
