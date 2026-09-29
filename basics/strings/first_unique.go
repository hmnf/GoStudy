package main

import (
	"slices"
	"strings"
)

func FirstUnique(s string) int {
	var deleted []string
	for i, l := range s {
		s = strings.Replace(s, string(l), "", 1)
		if !strings.Contains(s, string(l)) && !slices.Contains(deleted, string(l)) {
			return i
		}
		deleted = append(deleted, string(l))
	}

	return -1
}
