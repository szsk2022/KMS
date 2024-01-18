package main

import "fmt"

func main() {
	var a int
	fmt.Print("请输入一个数字：")
	fmt.Scanln(&a)
	fmt.Println("输出:", a)
}
