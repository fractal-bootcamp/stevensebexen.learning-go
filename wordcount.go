package main

import (
	"fmt"
	"strings"
)

func WordCounts(s string) map[string]int {
	m := make(map[string]int)
	fields := strings.Fields(s)
	for _, field := range fields {
		_, exists := m[field]
		if exists {
			m[field] += 1
		} else {
			m[field] = 1
		}
	}
	return m
}

func WordCount() {
	var test string = "Do you are you would you do do would you bananas?"
	counts := WordCounts(test)
	fmt.Println(counts)
}
