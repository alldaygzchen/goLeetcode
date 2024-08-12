package sliding_window

/*

s1="abb" s2="eidbbaooo"

l=0,r=0,need=len(s1),have=0

l=0,r=0 => {"e":1},have=0
l=0,r=1 => {"e":1,"i":1},have=0
l=0,r=2 => {"e":1,"i":1,"d":1},have=0
l=1,r=3 => {"b":1,"i":1,"d":1},have=0
l=2,r=4 => {"b":2,"d":1},have=1
l=3,r=5 => {"b":2,"a":1},have=2

s1="abb" s2="eidbbcabb"

l=0,r=0 => {"e":1},have=0
l=0,r=1 => {"e":1,"i":1},have=0
l=0,r=2 => {"e":1,"i":1,"d":1},have=0
l=1,r=3 => {"b":1,"i":1,"d":1},have=0
l=2,r=4 => {"b":2,"d":1},have=1
l=3,r=5 => {"b":2,"c":1},have=1
l=4,r=6 => {"b":1,"c":1,"a":1},have=1
l=5,r=7 => {"c":1,"a":1,"b":1},have=1
l=6,r=8 => {"a":1,"b":1,"b":1},have=1
*/

func checkInclusion(s1 string, s2 string) bool {


	l:=0
	r:=0
	need:=len(s1)
	have:=0
	s1Count := make(map[byte]int)
	s2Count := make(map[byte]int)

	if len(s1)>len(s2){
		return false
	}

	for i := 0; i< len(s1); i++{
        s1Count[s1[i]]++
    }

	for r<len(s2){
		s2Count[s2[r]]++

		if count, exists := s1Count[s2[r]]; exists {
			if count==s2Count[s2[r]]{
				have++
			}
		}

		if r+1==len(s1){
			if  count, exists := s1Count[s2[l]]; exists{
				if s2Count[s2[l]] == count {
                    have--
                }
			}
			s2Count[s2[1]]--
			if s2Count[s2[1]] == 0 {
				delete(s2Count,s2[1])
			}
			l++
		}

        if have == need {
            return true
        }

        r++
	}
	return false
}

