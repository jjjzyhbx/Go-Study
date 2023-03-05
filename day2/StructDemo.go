package main

import "fmt"

/**
这是一个关于go语言的结构体练习
*/

/**
在 Go 语言中，结构体是一种用户自定义的数据类型，
它可以由若干个任意类型的成员变量组成。
结构体可以用来表示一些复杂的数据结构，比如树、图、列表等，也可以用来表示一些实体，
比如人、汽车、房屋等。在 Go 语言中，结构体也是非常常用的数据类型之一。
定义结构体可以使用 type 关键字，语法如下：
*/

type person struct {
	name string
	age  int
}

// 结构体内嵌
type student struct {
	id  string "student id" //"student id" 这个字段是一个标签，
	per person "student is person"
}

// 成员函数
func (s student) show(_ int) {
	fmt.Println(s.id, s)
}

func main() {

	//例1
	var p person
	p.age = 22
	p.name = "宋"
	fmt.Println(p)

	var stu student
	stu.per.age = 18
	stu.per.name = "wag"
	stu.id = "260527496"
	fmt.Println(stu)
	//声明+初始化
	s := student{id: "260", per: person{"xiao", 18}}
	fmt.Println(s)
	s.show(22)

}
