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
}
