package tripletsum

import (
	"sort"
)

func FindTripletSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		pairsFound := findPairSum(nums[i+1:], -nums[i])
		result = append(result, pairsFound...)
	}

	return result
}

func findPairSum(sortedNumbers []int, target int) [][]int {
	l := 0
	r := len(sortedNumbers) - 1
	resp := [][]int{}

	for l < r {
		sum := sortedNumbers[l] + sortedNumbers[r]
		if sum == target {
			resp = append(resp, []int{sortedNumbers[l], sortedNumbers[r], -target})
			l++
			for l < r && sortedNumbers[l] == sortedNumbers[l-1] {
				l++
			}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return resp
}
