package main

import (
	"fmt"
)

func print1(c1 chan bool, c2 chan bool) {
	for i := 1; i < 101; i += 2 {
		c1 <- true
		fmt.Println(i)
		<-c2
	}
}
func print2(c1 chan bool, c2 chan bool, isDone chan bool) {
	for i := 2; i < 101; i += 2 {
		<-c1
		fmt.Println(i)
		c2 <- true
	}
	isDone <- true
}
func main() {
	c1 := make(chan bool)
	c2 := make(chan bool)
	isDone := make(chan bool)
	go print1(c1, c2)
	go print2(c1, c2, isDone)
	<-isDone
}
