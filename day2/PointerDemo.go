package main

import "fmt"

/*
*
这是一个go语言指针练习
*/
func main() {
	str := "hello"
	var p = &str
	*p = "world"
	fmt.Println(str)
	var a int = 10                         // 声明一个整型变量a
	var pa *int                            // 声明一个整型指针变量pa
	pa = &a                                // 将变量a的地址赋值给指针变量pa
	fmt.Printf("a的值为：%d\n", a)             // 输出a的值
	fmt.Printf("a的地址为：%p\n", &a)           // 输出a的地址
	fmt.Printf("pa的值为：%p\n", pa)           // 输出指针变量pa的值，即变量a的地址
	fmt.Printf("指针变量pa指向的变量的值为：%d\n", *pa) // 输出指针变量pa指向的变量的值，即变量a的值
}
