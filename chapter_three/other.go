package main

import (
	"bytes"
	"fmt"
)

/*
格式化动词
%v 原来的值
%T 变量的类型
%p 指针
%f 浮点数
%d 十进制
*/
func main() {
	// 字符串的拼接可以用+
	a := "01234"
	b := "5678"
	c := a + b
	fmt.Println(c)
	fmt.Println("-")

	// 如果拼接的字符串较长，可以使用字节缓冲的方式来进行
	trait := "kind of she "
	appearance := "gorgeous of she"
	var zero bytes.Buffer
	zero.WriteString(trait)
	zero.WriteString(appearance)
	fmt.Println(zero.String())

	str := "心之所向就是go with the flow"
	bytes1 := []byte(str)
	// 一个汉字占三个字节
	for i := 0; i <= 11; i++ {
		bytes1[i] = 'i'
	}
	fmt.Println(string(bytes1))

}
