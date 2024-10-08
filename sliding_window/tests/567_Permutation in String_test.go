package tests

import (
	"goLeetcode/sliding_window"
	"testing"
)


func Test_checkInclusion(t *testing.T) {
    input1 :="ab"
	input2 :="eidbaooo"
    expected := true
    result := sliding_window.CheckInclusion(input1,input2)
    if result == expected {
        t.Logf("Success: Expected %v and got %v", expected, result)
    } else {
        t.Errorf("Failure: Expected %v but got %v", expected, result)
    }
}
