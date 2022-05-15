package main

import (
	"fmt"
	"runtime"
)

func main() {
	if num := runtime.NumCPU(); num >= 2 {
		fmt.Printf("CPU数量是%d\n", num)
	}
	//	if语句的另一种写法

	for num1 := 1; num1 <= 10; num1 += 2 {
		fmt.Println("num1=", num1)
	}
	//	for循环语句

OuterLoop:
	for {
		i := 1
		for {
			if i < 5 {
				fmt.Println("退出for循环")
				break OuterLoop
			}
			i++
			fmt.Println("退出for外层循环")
		}
	}
	//	使用标签退出循环
}
