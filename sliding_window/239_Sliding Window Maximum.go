package sliding_window

/*
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]

1
1,3
1,3,-1=>3,-1
3,-1,-3=>-1,-3
-1,-3,5=>-3,5
-3,5,3=>5,3
5,3,6=>3,6
3,6,7=6,7

*/

func maxSlidingWindow(nums []int, k int) []int {
	queue:=make([]int,0)
	result:=make([]int,0)
	l:=0
	r:=0

	for r<len(nums){
		// decreasing or equivalent queue
		for len(queue)!=0 && nums[r]>queue[len(queue)-1]{

			queue=queue[:len(queue)-1]
		
		}
		
		queue = append(queue,nums[r])

        if r-l+1>=k{
            result = append(result,queue[0])
            if nums[l] == queue[0]{
			    queue=queue[1:]
			    l++
		    }
        }
		r++

	}
	return result
}





// func MaxSlidingWindow(nums []int, k int) []int {
// 	l:=0
// 	r:=0
// 	currMax := math.MinInt
// 	result := []int{}

// 	for r <len(nums){

//         if nums[r]>=currMax{
// 				currMax = nums[r]
// 		}

// 		if r+1>=k{
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
// 		// next loop
// 		r++
// 	}
// 	return result
// }