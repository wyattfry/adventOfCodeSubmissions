package common

import (
	"testing"
)

type extractIntTestCase struct {
	input string
	want  []int
}

func Test_ExtractInts(t *testing.T) {
	for _, tc := range []extractIntTestCase{
		{
			input: "testing: 1 2 test 34 test",
			want:  []int{1, 2, 34},
		},
		{
			input: "testing: 1 -2 test 34 test",
			want:  []int{1, -2, 34},
		},
	} {
		got := ExtractInts(tc.input)
		AssertEqual(tc.want, got, t)
	}
}
