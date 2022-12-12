package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2022/util"
)

func TestSolveCurrentDayWithTwist(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 29},
		{input: util.GetInput(), want: 418},
	}

	for _, tc := range tests {
		got := SolveCurrentDayWithTwist(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
