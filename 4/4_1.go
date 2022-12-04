package main

import (
	"strconv"
	"strings"
)

func SolveCurrentDay(input string) int {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		ranges := ParseLine(line)

		if ranges[0].Contains(ranges[1]) || ranges[1].Contains(ranges[0]) {
			count += 1
		}
	}

	return count
}

type Range struct {
	Start int
	End   int
}

func (r Range) Contains(other Range) bool {
	return other.Start >= r.Start && other.End <= r.End
}

func ParseLine(line string) [2]Range {
	ranges := strings.Split(line, ",")
	return [2]Range{
		ParseRange(ranges[0]),
		ParseRange(ranges[1]),
	}
}

func ParseRange(rangeStr string) Range {
	ends := strings.Split(rangeStr, "-")

	start, err := strconv.Atoi(ends[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(ends[1])
	if err != nil {
		panic(err)
	}

	return Range{
		Start: start,
		End:   end,
	}
}
