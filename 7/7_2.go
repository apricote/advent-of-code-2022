package main

import "math"

func SolveCurrentDayWithTwist(input string) int {
	fileTree := ParseInput(input)

	TOTAL_SPACE := 70_000_000
	UPDATE_SPACE := 30_000_000

	rootSize := fileTree.Size()
	availableSpace := TOTAL_SPACE - rootSize
	neededSpace := UPDATE_SPACE - availableSpace

	currentBestDirSize := math.MaxInt

	WalkTheTree(fileTree, func(d *Directory) {
		size := d.Size()

		if size > neededSpace && size < currentBestDirSize {
			currentBestDirSize = size
		}
	})

	return currentBestDirSize
}
