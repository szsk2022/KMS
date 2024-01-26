package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	update(arr)
}
func update(arr2 [4]int) {
	arr2[0] = 2
	fmt.Println("修改后的", arr2)
}
