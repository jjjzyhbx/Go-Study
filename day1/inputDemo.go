package main

import "fmt"

/**
在 Go 语言中，常用的输入输出语句是 fmt 包中的函数，其中常用的有：

fmt.Println()：输出一个或多个值并换行
fmt.Print()：输出一个或多个值不换行
fmt.Printf()：格式化输出
fmt.Scan()：从标准输入读取一个值
fmt.Scanln()：从标准输入读取一行数据
fmt.Scanf()：按照指定格式从标准输入读取数据
*/

func main() {
	var a string
	var b int
	var c bool
	var d float64
	var e complex128
	fmt.Scan(&a)

	fmt.Scanf("%d", &b)

	fmt.Scanln(&c)

	fmt.Scanln(&d)

	fmt.Scanln(&e)
	fmt.Println(a, b, c, d, e)
}
