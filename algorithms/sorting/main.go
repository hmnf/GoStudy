package main

func main() {
}

func arrEqual(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return false
		}
	}

	return true
}
