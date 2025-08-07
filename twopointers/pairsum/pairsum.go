package pairsum

func FindPairSum(sortedNumbers []int, target int) []int {
	l := 0
	r := len(sortedNumbers) - 1
	resp := []int{}

	for r < l {
		sum := sortedNumbers[r] + sortedNumbers[l]
		if sum > target {
			l++
		} else if sum < target {
			r--
		} else {
			resp = append(resp, r, l)
			break
		}
	}

	return resp
}
