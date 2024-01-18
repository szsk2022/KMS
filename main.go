package main

func main() {
	sum := 0
	for i := 1; i < 100000000000000000; i++ {
		sum = sum + i
		println(sum)
	}

}
