package main

import "fmt"

func main() {
	fmt.Print("请输入你的成绩：")
	var a int
	fmt.Scanln(&a)
	switch a {
	case 90:
		fmt.Print("A")
	case 80:
		fmt.Print("B")
	default:
		fmt.Print("D")

	}
}
