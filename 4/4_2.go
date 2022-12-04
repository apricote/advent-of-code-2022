package main

import "strings"

func SolveCurrentDayWithTwist(input string) int {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		ranges := ParseLine(line)

		if ranges[0].Overlaps(ranges[1]) || ranges[1].Overlaps(ranges[0]) {
			count += 1
		}
	}

	return count
}

func (r Range) Overlaps(other Range) bool {
	return other.Start >= r.Start && other.Start <= r.End || other.End >= r.Start && other.End <= r.End
}
