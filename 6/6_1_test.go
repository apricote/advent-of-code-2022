package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2022/util"
)

func TestSolveCurrentDay(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 7},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 5},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", want: 6},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 10},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 11},
		{input: util.GetInput(), want: 1275},
	}

	for _, tc := range tests {
		got := SolveCurrentDay(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
