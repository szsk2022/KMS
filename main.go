package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("默认：", arr)
	update(arr)
	fmt.Println("调用后的：", arr)
}
func update(arr2 []int) {
	fmt.Println("传递的：", arr2)
	arr2[0] = 2
	fmt.Println("修改后的", arr2)
}
