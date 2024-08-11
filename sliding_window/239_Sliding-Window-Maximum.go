package sliding_windowfunc

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {
        return []int{}
    }

    result := make([]int, 0, len(nums)-k+1)
    left, right := 0, 0
    maxIndex := 0

    for right < len(nums) {
        // Expand the window
        if right-left+1 <= k {
            if nums[right] >= nums[maxIndex] {
                maxIndex = right
            }
            right++
        }

        // Window is full, process it
        if right-left == k {
            result = append(result, nums[maxIndex])

            // If the maximum is leaving the window, find the new maximum
            if maxIndex == left {
                maxIndex = left + 1
                for i := left + 1; i < right; i++ {
                    if nums[i] >= nums[maxIndex] {
                        maxIndex = i
                    }
                }
            }

            left++
        }
    }

    return result
}