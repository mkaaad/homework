package main

import "fmt"

func main() {
	fmt.Println("欢迎使用Go语言计算器！")
	fmt.Println("请输入两个整数和一个操作符，进行四则运算")
	fmt.Println("输入exit退出程序")
	for true {
		var n, m, ans int
		var operator string
		var x string
		fmt.Println("请输入第一个整数")
		fmt.Scan(&n)
		fmt.Println("请输入操作符")
		fmt.Scan(&operator)
		fmt.Println("请输入第二个整数")
		fmt.Scan(&m)
		if operator == "/" && m == 0 {
			fmt.Println("被除数不能为零")
			continue
		}
		switch operator {
		case "+":
			ans = n + m
		case "-":
			ans = n - m
		case "*":
			ans = n * m
		case "/":
			ans = n / m
		default:
			continue
		}

		fmt.Printf("%d %s %d = %d", n, operator, m, ans)
		fmt.Println("是否继续？（exit退出）")
		fmt.Scan(&x)
		if x == "exit" {
			break
		}
	}
}
