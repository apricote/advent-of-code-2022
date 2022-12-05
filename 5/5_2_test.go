package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2022/util"
)

func TestSolveCurrentDayWithTwist(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: util.GetExampleInput(), want: "MCD"},
		{input: util.GetInput(), want: ""},
	}

	for _, tc := range tests {
		got := SolveCurrentDayWithTwist(tc.input)

		if tc.want != got {
			t.Errorf("Expected \"%s\" but got \"%s\"", tc.want, got)
		}
	}
}
