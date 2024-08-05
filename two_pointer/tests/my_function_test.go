package tests

import (
	"goLeetcode/two_pointer"
	"testing"
)

func TestExampleFunction(t *testing.T) {
    input := []int{1, 2, 3}
    expected := 1
    result := two_pointer.ExampleFunction(input)
    if result == expected {
        t.Logf("Success: Expected %d and got %d", expected, result)
    } else {
        t.Errorf("Failure: Expected %d but got %d", expected, result)
    }
}