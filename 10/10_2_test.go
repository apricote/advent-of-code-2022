package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2022/util"
	//"github.com/apricote/advent-of-code-2022/util"
)

const (
	part2ExampleOutput = `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`

	part2Output = `
###..####.#..#.###..###..#....#..#.###..
#..#.#....#..#.#..#.#..#.#....#..#.#..#.
#..#.###..####.#..#.#..#.#....#..#.###..
###..#....#..#.###..###..#....#..#.#..#.
#.#..#....#..#.#....#.#..#....#..#.#..#.
#..#.####.#..#.#....#..#.####..##..###..`  // REHPRLUB
)

func TestSolveCurrentDayWithTwist(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: util.GetExampleInput(), want: part2ExampleOutput},
		{input: util.GetInput(), want: part2Output},
	}

	for _, tc := range tests {
		got := SolveCurrentDayWithTwist(tc.input)

		if tc.want != got {
			t.Errorf("Expected:\n%s\n but got \n%s\n", tc.want, got)
		}
	}
}
