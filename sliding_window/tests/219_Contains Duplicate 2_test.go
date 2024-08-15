package tests

import (
	"goLeetcode/sliding_window"
	"testing"
)


func Test_ContainsNearbyDuplicate(t *testing.T) {
	input1 := []int{1,2,3,1,2,3} ///slice
	input2:=2
	expected := false
	result := sliding_window.ContainsNearbyDuplicate(input1, input2)
	if result == expected {
		t.Logf("Success: Expected %v and got %v", expected, result)
	} else {
		t.Errorf("Failure: Expected %v but got %v", expected, result)
	}
}