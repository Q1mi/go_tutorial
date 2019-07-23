package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	r1 := calc(100, 200, add)
	fmt.Println(r1)
	r2 := calc(100, 200, sub)
	fmt.Println(r2)
}
