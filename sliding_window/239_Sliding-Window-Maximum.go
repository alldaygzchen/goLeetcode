package sliding_windowfunc

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	l:=0
	r:=0
	currMax := nums[r]
	result := []int{}


	for r <len(nums){

		if r-l+1<k{

			if nums[r]>=currMax{
				currMax = nums[r]
			}
		}else{
			if nums[r]>=currMax{
				currMax = nums[r]
				result = append(result, currMax)
			}
			if nums[l]==currMax{
				currMax = nums[l+1]
				for i:=l; i<r+1;i++{
					if nums[i]>=currMax{
						currMax = nums[i]
					}
				}
			}
			result = append(result, currMax)

			l++
		}
		r++
		fmt.Println("result",result)
		fmt.Println("l",l,"r",r)
	}
	return result
}