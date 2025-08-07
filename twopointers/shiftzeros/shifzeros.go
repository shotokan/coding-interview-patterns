package shiftzeros

func ShiftZerosToEnd(values []int) []int {
	left, right := 0, 0

	for right < len(values) {
		if values[right] != 0 {
			values[left], values[right] = values[right], values[left]
			left++

		}

		right++
	}
	return values
}
