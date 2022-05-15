package main

import (
	"fmt"
	"reflect"
)

/*
整型：int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64
浮点型：float32 float64
字符串：string
字符 uint8 代表了ascii码的一个字符 rune 代表了utf格式的一个字符，可以是中文，本质是int32类型
转义字符
\n 换行符
\r 回车符
\t 制表符
\' 单引号
\" 双引号
\\ 反斜杠
布尔型 bool
false true

*/
func main() {
	var1 := "20.0"
	fmt.Printf("var1的类型是%v", reflect.TypeOf(var1))

	var2 := `
one
two 
three`
	fmt.Println(var2)
	//	反引号可以对内容直接输出

	chinese := '非'
	fmt.Println(chinese) // 直接打印输出的是码值
	fmt.Printf("%c\n", chinese)

	num1 := 80
	num2 := int32(num1)
	fmt.Println(num2, "类型是", reflect.TypeOf(num2))
}
