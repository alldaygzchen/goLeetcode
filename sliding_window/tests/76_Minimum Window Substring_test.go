package tests

import (
	"goLeetcode/sliding_window"
	"testing"
)


func Test_minWindow(t *testing.T) {
    input1 :="ADOBECODEBANC"
	input2 :="ABC"
    expected := "BANC"
    result := sliding_window.MinWindow(input1,input2)
    if result == expected {
        t.Logf("Success: Expected %v and got %v", expected, result)
    } else {
        t.Errorf("Failure: Expected %v but got %v", expected, result)
    }
}
