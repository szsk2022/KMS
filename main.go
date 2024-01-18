package main

import "fmt"

func main() {
	fmt.Println("Hello,World!")
	fmt.Println(add(1, 2))
}
func add(a, b int) int {
	c := a + b
	return c
}
