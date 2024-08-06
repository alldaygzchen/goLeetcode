package two_pointer

func Trap(height []int) int {

	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0

	for left <= right {

		if height[left] <= height[right] {
			if height[left] <= maxLeft {
				res += maxLeft - height[left]
			} else {
				maxLeft = height[left]
			}
			left++

		} else {
			if height[right] <= maxRight {
				res += maxRight - height[right]
			} else {
				maxRight = height[right]
			}
			right--
		}

	}
	return res

}