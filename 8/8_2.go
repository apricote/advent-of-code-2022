package main

func SolveCurrentDayWithTwist(input string) int {
	trees := ParseInput(input)

	best := -1
	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[0]); x++ {
			scenicScore := GetScenicScore(trees, y, x)

			if scenicScore > best {
				best = scenicScore
			}
		}
	}

	return best
}

func GetScenicScore(trees [][]int, y, x int) int {
	height := trees[y][x]

	treesRight := 0
	for xOther := x + 1; xOther < len(trees[y]); xOther++ {
		treesRight++

		if trees[y][xOther] >= height {
			break
		}
	}

	treesLeft := 0
	for xOther := x - 1; xOther >= 0; xOther-- {
		treesLeft++

		if trees[y][xOther] >= height {
			break
		}
	}

	treesBottom := 0
	for yOther := y + 1; yOther < len(trees); yOther++ {
		treesBottom++

		if trees[yOther][x] >= height {
			break
		}
	}

	treesTop := 0
	for yOther := y - 1; yOther >= 0; yOther-- {
		treesTop++

		if trees[yOther][x] >= height {
			break
		}
	}

	return treesRight * treesLeft * treesBottom * treesTop

}
