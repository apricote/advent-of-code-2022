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
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 19},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 23},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", want: 23},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 29},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 26},
		{input: util.GetInput(), want: 0},
	}

	for _, tc := range tests {
		got := SolveCurrentDayWithTwist(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
