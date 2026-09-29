package main

func FirstUnique(s string) int {
	count := make(map[rune]int)

	for _, l := range s {
		count[l]++
	}

	for i, l := range s {
		if count[l] == 1 {
			return i
		}
	}

	return -1
}
