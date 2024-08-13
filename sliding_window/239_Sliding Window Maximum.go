package sliding_window

// time exceed
// func maxSlidingWindow(nums []int, k int) []int {
// 	l:=0
// 	r:=0
// 	currMax := math.MinInt
// 	result := []int{}

// 	for r <len(nums){

//         if nums[r]>=currMax{
// 				currMax = nums[r]
// 		}

// 		if r-l+1==k{
//             result = append(result, currMax)
//             if nums[l]==currMax{
//                 currMax = math.MinInt
//                 for i:=l+1; i<r+1;i++{
//                     if nums[i]>=currMax{
//                         currMax = nums[i]
//                     }
//                 }

//             }
//             l++
// 		}
// 		r++
// 	}
// 	return result
// }