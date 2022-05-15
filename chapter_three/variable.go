package main

import "fmt"

/*
var 变量名 变量类型
var num int

go语言在声明变量时，会对每个变量的内存区域进行初始化，即每个变量都有对应的默认值
切片、映射、函数、指针变量默认为nil，相当于其它语言的null None

*/
func main() {
	var name string
	name, age := "zero", 19
	fmt.Println(name, "is", age)
	//	应至少有一个新的变量出现，这样编译才不会报错

	num1 := 11
	num2 := 22
	num3 := 33
	num1, num2, num3 = num2, num3, num1
	fmt.Println(num1, num2, num3)
	//	多重赋值时，按照从左到右的顺序赋值

}

func ReturnData() (int, int) {
	return 10, 20
	//	匿名变量，就是不需要接受值的变量
	//	匿名变量既不占用命名空间，也不分配系统内存
}
