package sliding_window

import (
	"math"
)

/*
s = "ADOBECODEBANC", t = "ABC"

A
AD
ADO
ADOB
ADOBE
DOBEC
DOBECO
DOBECOD
DOBECODE
DOBECODEB
ODEBA
ODEBAN
ANC
*/

func MinWindow(s string, t string) string {

	l:=0
	r:=0
	output := ""
	minLength := math.MaxInt
	countT:= make(map[byte]int)
	for i:=0;i<len(t);i++{
		countT[t[i]]++
	}
	countS:= make(map[byte]int)
	have := 0
	need := len(countT)

	for r<len(s){
		countS[s[r]]++
		if count,exist:=countT[s[r]];exist && count == countS[s[r]]{
			have++
		}

		for have==need {
			if r-l+1<minLength {
				minLength = r-l+1
				output = s[l:r+1]
			}
			countS[s[l]]--
			if count,exist:=countT[s[l]];exist && count == countS[s[l]]+1{
				have--
			}
			l++
		}
		// fmt.Println(s[l:r+1])
		// next loop
		r++
	}
	return output
}