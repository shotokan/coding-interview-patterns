package nextlexicographicalsequence

func NextLexicographicalSequence(input string) string {
	runes := []rune(input)

	pivot := len(runes) - 2

	for pivot >= 0 && runes[pivot] >= runes[pivot+1] {
		pivot--
	}

	if pivot < 0 {
		reverseSubstring(runes, 0)
		return string(runes)
	}

	rightMostSuccesor := len(runes) - 1

	for runes[rightMostSuccesor] <= runes[pivot] {
		rightMostSuccesor--
	}

	runes[pivot], runes[rightMostSuccesor] = runes[rightMostSuccesor], runes[pivot]
	reverseSubstring(runes, pivot+1)

	return string(runes)
}

func reverseSubstring(runes []rune, start int) {
	end := len(runes) - 1
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}
