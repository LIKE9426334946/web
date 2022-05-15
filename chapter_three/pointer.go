package main

import "fmt"

/*
指针是一种地址值，这个地址值代表着计算机内存空间中的某个位置，指针变量就是存放地址值的变量
指针变量未指向任何地址，所以指针值为nil
使用操作符"&"取变量地址
ptr1=&"kind"  不可以这样写，&后面要跟一个变量
指针也有自己的内存地址
*ptr 来获取指针指向的值
new()函数可以为指针新分配地址并初始化值
*/

func main() {
	var ptr1 *string // 声明类型为指针
	fmt.Println(ptr1)

	str1 := "kind"
	ptr1 = &str1
	fmt.Println("ptr1指向的值是", *ptr1)

	ptr1 = new(string)
	fmt.Printf("new()之后的ptr1为%q\n", *ptr1)
}
