package sliding_window

import "math"

/*

Input: s = "abcabcbb"
Output: 3

a
ab
abc
abca=>bca
bcab=>cab
cabc=>abc
abcb=>cb
cbb=>b

*/

func LengthOfLongestSubstring(s string) int {

	l:=0
	r:=0
	maxLength := math.MinInt
	countChar:=make(map[byte]int)

	if len(s)==0{
		return 0
	}

	for r<len(s){

		countChar[s[r]]++

		for countChar[s[r]]>1{
			countChar[s[l]]--
			l++
		}

		if r-l+1>maxLength{
			maxLength = r-l+1
		}

		r++
	}
	return maxLength


}