package sliding_window

/*

s1="abb" s2="eidbbcabb"

e
ei
eid
idb
dbb
bbc
bca
cab
abb

*/

func CheckInclusion(s1 string, s2 string) bool {



	if len(s1)>len(s2){
		return false
	}

	l:=0
	r:=0

	s1Count := make(map[byte]int)
	s2Count := make(map[byte]int)

	for i := 0; i< len(s1); i++{
        s1Count[s1[i]]++
    }

	need:=len(s1Count)
	have:=0

	for r<len(s2){
		s2Count[s2[r]]++

		if count, exists := s1Count[s2[r]]; exists {
			if count==s2Count[s2[r]]{
				have++
			}
		}

		if r+1>len(s1){
			if  count, exists := s1Count[s2[l]]; exists{
				if s2Count[s2[l]] == count {
                    have--
                }
			}
			s2Count[s2[l]]--
			l++
		}
        r++
		if have == need {
            return true
        }

	}
	return false
}

