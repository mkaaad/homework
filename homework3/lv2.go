package main

import (
	"fmt"
	"time"
)

// 定义全局变量
var ch chan int               // 用于goroutine间通信
var startTime time.Time       // 计时器开始时间
var elapsedTime time.Duration //计时器已运行时间
var running bool              // 计时器是否在运行
var myflag int

// Timer 函数，接收通道中的命令并执行相应的操作

func Timer(ch chan int) {
	for {
		switch <-ch {
		case 0: // 重置计时器
			startTime = time.Time{}
			elapsedTime = 0
			running = false
			fmt.Println("计时器已重置")
		case 1: // 开始计时
			if !running {
				startTime = time.Now()
				running = true
				fmt.Println("计时器已开始")
			} else {
				fmt.Println("计时器已在运行")
			}
		case 2: // 暂停/继续计时
			if running {
				elapsedTime += time.Since(startTime)
				running = false
				fmt.Println("计时器已暂停，已运行时间:", elapsedTime)
			} else {
				startTime = time.Now()
				running = true
				fmt.Println("计时器已继续")
			}
		case 3: //记录多人时间功能
			if running {
				flagTime := time.Since(startTime)
				myflag++
				fmt.Printf("第%v次计时，时间为%v\n", myflag, flagTime)
			}
		}
	}
}

// Input 函数，接收用户输入的命令并发送到通道 func Input(ch chan int) {
func Input(ch chan int) {
	for {
		var cmd int
		fmt.Scan(&cmd)
		ch <- cmd
	}
}

func main() {
	fmt.Println("请输入命令:0-重置计时器，1-开始计时，2-暂停/继续计时，3-添加flag")
	// 初始化通道
	ch = make(chan int)
	// 启动 Timer goroutine
	go Timer(ch)
	// 启动 Input goroutine
	go Input(ch)
	// 主goroutine 等待
	select {}
}
