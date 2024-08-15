package sliding_window

/*

arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4

2
2,2
2,2,2=>2,2
2,2,2=>2,2
2,2,5=>2,5
2,5,5=>5,5
5,5,5=>5,5
5,5,8=>5,8
*/



func numOfSubarrays(arr []int, k int, threshold int) int {

	count:=0
	l:=0
	r:=0
	total:=0



	for r<len(arr){

		if r-l+1==k{
			total=0
			for _ ,v :=range arr[l:r+1]{
				total=total+v
			}

			if total >= (threshold*k){
				count++
			}

			l++
		}

		r++

	}

	return count

}