package main

import "strings"

func reverse(str string) string {
	var j int
	var reverseStr string
	parts := strings.Fields(str)
	j = len(parts) - 1
	for j >= 0 {
		reverseStr = reverseStr + " " + parts[j]
		j--
	}
	return reverseStr[1:]
}
