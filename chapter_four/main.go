package main

import "fmt"

func main() {
	const a float64 = 3.1415
	//	显式定义

	const b = "hello"
	//	隐式定义

	const (
		zero = iota
		one
		two
	)
	fmt.Println("zero", zero, "one", one, "two", two)
	//	常量枚举

}
