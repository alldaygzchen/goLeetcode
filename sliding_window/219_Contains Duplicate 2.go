package sliding_window

/*
Input: nums = [1,2,3,1,2,3], k = 2
Output: false

1
1,2
1,2,3=>2,3
2,3,1=>3,1
3,1,2=<1,2
1,2,3=>2,3

*/
func ContainsNearbyDuplicate(nums []int, k int) bool {

	countMap:=make(map[int]int)
	l:=0
	r:=0

	for r <len(nums){

		countMap[nums[r]]++

		if countMap[nums[r]]>1{
			return true
		}

		if r-l+1 == k{
			countMap[nums[l]]--
			l++
		}
		r++
	}
	return false


}