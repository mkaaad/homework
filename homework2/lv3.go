package main

import (
	"errors"
	"fmt"
	"time"
)

type Task interface {
	Execute() error
}
type PrintTask struct {
	Message string
}
type CalculationTask struct {
	A int
	B int
}
type SleepTask struct {
	Duration int
}

func (p PrintTask) Execute() error {
	fmt.Println(p.Message)
	return nil
}
func (c CalculationTask) Execute() error {
	fmt.Println(c.A + c.B)
	return nil
}
func (s SleepTask) Execute() error {
	time.Sleep(time.Duration(s.Duration) * time.Second)
	return nil
}

type Scheduler struct {
	Tasks []Task
}

func (s *Scheduler) AddTask(task Task) error {
	s.Tasks = append(s.Tasks, task)
	return nil
}
func (s Scheduler) RunAll() error {
	if len(s.Tasks) != 0 {
		for _, task := range s.Tasks {
			task.Execute()
		}
		return nil
	} else {
		return errors.New("task list is empty")
	}
}
func main() {
	var x Scheduler
	p := PrintTask{"It's a task"}
	s := SleepTask{3}
	c := CalculationTask{5, 8}
	x.AddTask(p)
	x.AddTask(s)
	x.AddTask(c)
	text := x.RunAll()
	if text != nil {
		fmt.Println("Error!", text)
	}
	var y Scheduler
	text = y.RunAll()
	if text != nil {
		fmt.Println("Error!", text)
	}
}
