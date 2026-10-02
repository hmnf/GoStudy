package main

func main() {
}

func isValid(s string) bool {
	sk := make(map[rune]rune)
	sk['['] = ']'
	sk['('] = ')'
	sk['{'] = '}'
	var stack []int
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] == ']' || r[i] == '}' || r[i] == ')' {
			if sk[r[stack[len(stack)-1]]] == r[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, i)
		}
	}

	return true
}
