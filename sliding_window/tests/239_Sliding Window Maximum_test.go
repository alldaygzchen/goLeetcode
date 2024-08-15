package tests

import (
	"goLeetcode/sliding_window"
	"reflect"
	"testing"
)

func Test_MaxSlidingWindow(t *testing.T) {
	input1 := []int{1, 3, -1, -3, 5, 3, 6, 7} ///slice
	input2:=3
	expected := []int{3, 3, 5, 5, 6, 7}
	result := sliding_window.MaxSlidingWindow(input1, input2)
	if reflect.DeepEqual(result, expected) {
		t.Logf("Success: Expected %v and got %v", expected, result)
	} else {
		t.Errorf("Failure: Expected %v but got %v", expected, result)
	}
}