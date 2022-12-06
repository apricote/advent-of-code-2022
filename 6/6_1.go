package main

type Set map[string]struct{}

func (s Set) Add(element string) {
	s[element] = struct{}{}
}

func (s Set) Has(element string) bool {
	_, ok := s[element]

	return ok
}

func SolveCurrentDay(input string) int {
	return GetXUniquePosition(input, 4)
}

func GetXUniquePosition(input string, uniqueChars int) int {
out:
	for i := uniqueChars - 1; i < len(input); i++ {
		charMap := Set{}

		for j := 0; j < uniqueChars; j++ {
			currentChar := string(input[i-j])

			if charMap.Has(currentChar) {
				continue out //outer loop
			}

			charMap.Add(currentChar)
		}

		return i + 1
	}

	return 0
}
