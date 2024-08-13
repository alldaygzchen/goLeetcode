package gointro

import "fmt"

func check() {

	testMap := make(map[byte]int)
	testMap['a'] = 1
	fmt.Println("testMap",testMap)
	testMap[97] = 2
	fmt.Println("testMap",testMap)
	testMap[97] = 3
	fmt.Println("testMap",testMap)
}