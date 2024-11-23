package main

import (
	"fmt"
)

var goods int
var lock chan bool
func mutex(){
	lock<-true
}
func unmutex(){
	<-lock
}

func customer(buy chan string) {
	if <-buy == "buy goods" {
		mutex()
		defer unmutex()
		if goods > 0 {
			goods--
			fmt.Println("Buy sccessflly, goods remain:", goods)
		} else {
			fmt.Println("Buying failed: no goods remain")
		}
	}
}
func producer(produce chan string) {
	mutex()
	defer unmutex()
	if <-produce == "produce" {
		goods++
		fmt.Println("Produce sccessflly, goods remain:", goods)
	}
}
func main() {
	buy := make(chan string, 100)
	produce := make(chan string, 100)
	lock=make(chan bool,1)
	for i := 0; i < 7; i++ {
		produce <- "produce"
		go producer(produce)
	}
	for i := 0; i < 5; i++ {
		buy <- "buy goods"
		go customer(buy)
	}
	for len(buy)!=0 || len(produce)!=0{}
}
