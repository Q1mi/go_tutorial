package main

import "fmt"

func main() {
	// var a *int
	// a = new(int)
	// fmt.Println(*a) // 0
	// *a = 100
	// fmt.Println(*a) // 100

	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
