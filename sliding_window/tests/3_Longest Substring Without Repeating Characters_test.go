package tests

import (
	"goLeetcode/sliding_window"
	"testing"
)


func Test_LengthOfLongestSubstring(t *testing.T) {
    input1 :="abcabcbb"
    expected := 3
    result := sliding_window.LengthOfLongestSubstring(input1)
    if result == expected {
        t.Logf("Success: Expected %v and got %v", expected, result)
    } else {
        t.Errorf("Failure: Expected %v but got %v", expected, result)
    }
}
