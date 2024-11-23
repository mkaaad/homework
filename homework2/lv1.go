package main

import "fmt"

func main() {
	type Student struct {
		Name  string
		Age   int
		Score float32
	}
	var arr []Student
	arr = append(arr, Student{"Alen", 16, 86})
	arr = append(arr, Student{"Tom", 18, 95})
	for i, s := range arr {
		fmt.Println("Student", i+1)
		fmt.Println(" Name=", s.Name)
		fmt.Println(" Age=", s.Age)
		fmt.Println(" Score=", s.Score)
	}
}
