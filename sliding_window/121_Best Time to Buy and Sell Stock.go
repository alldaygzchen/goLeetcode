package sliding_window

import "math"

/*
prices = [7,2,5,3,6,4,1]
answer = 5

7
7,2=>2
2,5
2,5,3
2,5,2,6
2,5,2,6,4
2,5,2,6,4,1=>1
*/

func MaxProfit(prices []int) int {

	l:=0
	r:=0
	profixMax := math.MinInt
	// minPrice := math.MaxInt

	for r<len(prices){

		if prices[l]>prices[r]{
			l=r
		}else{
			if (prices[r]-prices[l])>profixMax{
				profixMax=prices[r]-prices[l]
			}
		}
		
		r++
	}
	return profixMax
}