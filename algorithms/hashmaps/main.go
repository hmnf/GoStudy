package main

func main() {
}

func isAnagram(s string, t string) bool {
	count := make(map[rune]int)

	for _, l := range s {
		count[l]++
	}

	for _, l := range t {
		count[l]--
		if count[l] == 0 {
			delete(count, l)
		}
	}

	if len(count) == 0 {
		return true
	}

	return false
}
