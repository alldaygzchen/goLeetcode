package sliding_window

import "math"

/*
Input: s = "AABABBA", k = 1
Output: 4

A
AA
AAB
AABA
AABAB => AB
ABB
ABBA=>BBA
*/


func CharacterReplacement(s string, k int) int {

	l:=0
	r:=0
	maxFreq:= 0
	maxLength:= math.MinInt
	countMap:=make(map[byte]int)

	for r<len(s){
		countMap[s[r]]++
		if countMap[s[r]]>maxFreq{
			maxFreq=countMap[s[r]]
		}

		if (r-l+1)-maxFreq>k{
			countMap[s[l]]--
			if maxFreq<countMap[s[l]]{
				maxFreq=countMap[s[l]]
			}
			l++
		}

		if (r-l+1)>maxLength{
			maxLength =r-l+1
		}

		r++
	}

	return maxLength
}
