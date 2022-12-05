package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2022/util"
)

func TestSolveCurrentDay(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: util.GetExampleInput(), want: "CMZ"},
		{input: util.GetInput(), want: "ZSQVCCJLL"},
	}

	for _, tc := range tests {
		got := SolveCurrentDay(tc.input)

		if tc.want != got {
			t.Errorf("Expected \"%s\" but got \"%s\"", tc.want, got)
		}
	}
}
