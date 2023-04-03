package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	sArr := strings.Fields(s)

	for i := range sArr {
		v, exists := m[sArr[i]]
		if exists {
			m[sArr[i]] = v + 1
		} else {
			m[sArr[i]] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
