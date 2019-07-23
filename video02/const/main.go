package main

import "fmt"

//常量
// const pi = 3.1415
// const e = 2.7

// const (
// 	pi = 3.1415
// 	e  = 2.7
// )

// const (
// 	n1 = 10
// 	n2
// 	n3
// )

const (
	n1 = iota //0
	n2 = iota //1
	n3 = 100  //100
	n4 = iota //3
)

const n5 = iota //0

const (
	_  = iota
	KB = 1 << (10 * iota) //1<<10
	MB = 1 << (10 * iota) //1<<20
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2//iota=0, 1,2
	c, d = iota+1,iota+2//iota=1,2,3
	e, f//iota=2,3,4
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
