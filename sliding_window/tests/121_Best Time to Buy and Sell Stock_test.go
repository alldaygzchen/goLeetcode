package tests

import (
	"goLeetcode/sliding_window"
	"testing"
)


func Test_maxProfit(t *testing.T) {
    input1 := []int{7,1,5,3,6,4} ///slice
    expected := 5
    result := sliding_window.MaxProfit(input1)
    if result == expected {
        t.Logf("Success: Expected %v and got %v", expected, result)
    } else {
        t.Errorf("Failure: Expected %v but got %v", expected, result)
    }
}
