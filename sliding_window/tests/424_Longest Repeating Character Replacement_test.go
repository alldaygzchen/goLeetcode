package tests

import (
	"goLeetcode/sliding_window"
	"testing"
)


func Test_CharacterReplacement(t *testing.T) {
    input1 := "ABAB"
    input2 := 2
	expected := 4
    result := sliding_window.CharacterReplacement(input1,input2)
    if result == expected {
        t.Logf("Success: Expected %v and got %v", expected, result)
    } else {
        t.Errorf("Failure: Expected %v but got %v", expected, result)
    }
}
