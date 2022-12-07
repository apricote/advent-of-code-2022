package main

import (
	"strconv"
	"strings"
)

func SolveCurrentDay(input string) int {
	fileTree := ParseInput(input)

	totalSize := 0

	WalkTheTree(fileTree, func(d *Directory) {
		currentDirSize := d.Size()
		if currentDirSize <= 100_000 {
			totalSize += currentDirSize
		}
	})

	return totalSize
}

func WalkTheTree(fileTree *Directory, actor func(*Directory)) {
	remainingDirs := []*Directory{fileTree}

	for len(remainingDirs) > 0 {
		currentDir := remainingDirs[0]
		remainingDirs = remainingDirs[1:]

		for _, subDir := range currentDir.Subdirectories {
			remainingDirs = append(remainingDirs, subDir)
		}

		actor(currentDir)
	}
}

type Directory struct {
	Files          []int
	Subdirectories map[string]*Directory
	Parent         *Directory
}

func NewDirectory(parent *Directory) *Directory {
	return &Directory{
		Subdirectories: map[string]*Directory{},
		Parent:         parent,
	}
}

func (d Directory) Size() int {
	size := 0

	for _, files := range d.Files {
		size += files
	}

	for _, dir := range d.Subdirectories {
		size += dir.Size()
	}

	return size
}

func ParseInput(input string) *Directory {
	rootDir := NewDirectory(nil)
	var currentDir *Directory

	for _, line := range strings.Split(input, "\n") {
		// Command Input
		// - cd
		if strings.HasPrefix(line, "$ cd ") {
			targetDir := strings.TrimPrefix(line, "$ cd ")
			// root dir special case
			switch targetDir {
			case "/":
				currentDir = rootDir
			case "..":
				currentDir = currentDir.Parent
			default:
				currentDir = currentDir.Subdirectories[targetDir]
			}

			continue
		}

		// - ls
		if strings.HasPrefix(line, "$ ls") {
			// nothing needs to be done

			continue
		}

		// ls output
		// - dir
		if strings.HasPrefix(line, "dir ") {
			dirName := strings.TrimPrefix(line, "dir ")
			currentDir.Subdirectories[dirName] = NewDirectory(currentDir)

			continue
		}
		// - file
		// nothing else matched, must be file
		parts := strings.Split(line, " ")
		sizeStr := parts[0]
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			panic(err)
		}
		currentDir.Files = append(currentDir.Files, size)
	}

	return rootDir
}
