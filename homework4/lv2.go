package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("events.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var before, after time.Duration
	var beforeEvent, afterEvent string
	now := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		eventTime, _ := time.Parse("2006-01-02", text[0])

		if now.Sub(eventTime) > 0 {
			before = now.Sub(eventTime)
			beforeEvent = text[1]
		} else {
			after = now.Sub(eventTime)
			afterEvent = text[1]
			break
		}
	}
	if (before + after) < 0 {
		fmt.Println(beforeEvent)
	} else if (before + after) > 0 {
		fmt.Println(afterEvent)
	} else {
		fmt.Println(beforeEvent, afterEvent)
	}
}
