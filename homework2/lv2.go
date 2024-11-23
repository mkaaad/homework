package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name   string
	Age    int
	Scores []int
}
type Classroom struct {
	Classname string
	sigle     []*Student
}

func AddStudent(c *Classroom, s *Student) {
	c.sigle = append(c.sigle, s)
}
func UpdateScore(s *Student, score int) {
	s.Scores = append(s.Scores, score)
}
func CalculateAverage(s *Student) float64 {
	sum := 0
	for _, score := range s.Scores {
		sum += score
	}
	return float64(sum) / float64(len(s.Scores))
}
func main() {
	class1 := &Classroom{}
	stu1 := &Student{
		Name: "Sam",
		Age:  16,
	}
	stu2 := &Student{
		Name: "Tom",
		Age:  17,
	}
	stu3 := &Student{
		Name: "Jack",
		Age:  15,
	}
	AddStudent(class1, stu1)
	AddStudent(class1, stu2)
	AddStudent(class1, stu3)
	for i := 1; i < 8; i++ {
		UpdateScore(stu1, rand.Intn(41)+60)
		UpdateScore(stu2, rand.Intn(41)+60)
		UpdateScore(stu3, rand.Intn(41)+60)
	}
	fmt.Println(CalculateAverage(stu1))
	fmt.Println(CalculateAverage(stu2))
	fmt.Println(CalculateAverage(stu3))
}
