package main

/**
这是一个循环控制语句的详解
*/
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/**
	for 循环
	for 循环是 Go 语言中最基本的循环控制语句，它用于重复执行某段代码块，
	直到指定的条件不再满足。for 循环有以下三种形式：
	其中，init 语句在循环开始前执行，condition 表示循环条件，post 语句在每次循环结束后执行。
	for init; condition; post { }
	for condition { }
	for { }
	*/
	// for 循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	rand.Seed(time.Now().UnixNano())
	//等价与while循环
	for rand.Intn(21)-10 < 0 {
		i := 1
		fmt.Println("第", i, "次", rand.Intn(21)-10)
		i++
	}

	// break 语句
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// continue 语句
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// goto 语句
	i := 0
Loop:
	if i < 10 {
		fmt.Println(i)
		i++
		goto Loop
	}

	fmt.Println("循环进阶----------------")
	//for-range 循环
	//for-range 循环可以用来遍历数组、切片、字符串、映射等集合类型。for-range 循环的语法格式如下
	//for key, value := range collection {
	//	// 循环体
	//}
	//其中，key 和 value 分别表示当前迭代的元素的索引和值，
	//collection 表示要迭代的集合类型。在循环体内部，可以使用 break 和 continue 关键字来控制循环的执行流程。

	arr := []int{1, 2, 6, 4, 8, 99, 10}
	for key, value := range arr {
		fmt.Println("arr[", key, "]=", value)
	}
	// 遍历字符串
	str := "Hello, world!"
	for i, c := range str {
		fmt.Println(i, string(c))
	}

}
