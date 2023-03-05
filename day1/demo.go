package main

/**
这是一个基础类型的文件
*/
//导入语句 import Name "packageName"  Name可以省略
import fmt1 "fmt"

func main() {
	//基础定义语句
	fmt1.Println("基础定义语句")
	var a int //默认0
	var b int = 1
	var c = 1
	d := 1
	fmt1.Println(a, b, c, d)
	//简单类型
	fmt1.Println("简单类型")
	var a1 int8 = 1
	var b1 byte = 'a'
	var c1 float64 = 1.3
	var str = "sdh"
	ok := true

	fmt1.Println(a1, b1, c1, str, ok)
	//字符串
	fmt1.Println("字符串")
	str1 := "go study"
	str2 := "java study"
	runeArr := []rune(str2)
	fmt1.Println(str1[1], str2[2], runeArr[2], string(runeArr[2]))

	//数组声明
	var arr1 [5]int
	var arr2 [5][4]int
	//声明+初始化
	arr3 := []int{1, 5, 1, 5, 2, 5, 52, 5, 2, 2}
	fmt1.Println("数组声明")
	fmt1.Println(arr1, arr2, arr3)
	fmt1.Println("长度", len(arr3))
	//动态数组，切片
	slice1 := make([]float64, 0)   //长度为0的切片
	slice2 := make([]string, 3, 5) //长度为3容量为5的切片
	fmt1.Println(slice1, slice2)

	fmt1.Println("添加")
	slice2 = append(slice2, "element1", "element2", "wiudh")
	fmt1.Println(slice2, "长度", len(slice2), "容量", cap(slice2))

	fmt1.Println("子切片")
	fmt1.Println("0-4", slice2[1:4], "长度", len(slice2[1:4]), "容量", cap(slice2[1:4]))
	fmt1.Println("2-5", slice2[2:5], "长度", len(slice2[2:5]), "容量", cap(slice2[2:5]))
	//子切片容量为底层容量-子切片起始容量

	//map 键值对
	fmt1.Println("map 键值对")
	m1 := make(map[string]int)
	m2 := map[string]int{
		"go":   1,
		"java": 2,
	}
	fmt1.Println(m1, m2)
	fmt1.Println("添加")
	m2["c++"] = 3
	m2["str"] = 22
	fmt1.Println(m2)

	fmt1.Println("修改")
	m2["c++"] = 23
	fmt1.Println(m2)

	fmt1.Println("删除")
	delete(m2, "str")
	fmt1.Println(m2)

	fmt1.Println("查找")
	value, exists := m2["go"]
	if exists {
		fmt1.Println("go:", value)
	} else {
		fmt1.Println("go not found")
	}
	fmt1.Println(m2)

}
