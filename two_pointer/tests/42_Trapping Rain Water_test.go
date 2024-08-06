package tests

import (
	"goLeetcode/two_pointer"
	"testing"
)


func Test_Problem42(t *testing.T) {
    input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
    expected := 6
    result := two_pointer.Trap(input)
    if result == expected {
        t.Logf("Success: Expected %d and got %d", expected, result)
    } else {
        t.Errorf("Failure: Expected %d but got %d", expected, result)
    }
}
