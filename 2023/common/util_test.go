package common

import (
	"reflect"
	"slices"
	"testing"
)

func Test_ExtractInts(t *testing.T) {
	input := "testing: 1 2 test 34 test"
	want := []int{1, 2, 34}
	result := ExtractInts(input)
	slices.Sort(want)
	slices.Sort(result)
	if !reflect.DeepEqual(want, result) {
		t.Error("ExtractInts() =", result, "but wanted", want)
	}
}
