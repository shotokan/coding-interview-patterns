package largestcontainer

import "math"

func LargestContainer(heights []int) float64 {
	left, right := 0, len(heights)-1
	max := float64(0)

	for left < right {
		min := math.Min(float64(heights[left]), float64(heights[right]))
		result := min * (float64(right) - float64(left))
		max = math.Max(max, result)
		if heights[left] == heights[right] {
			left++
			right--
		} else if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return max
}
