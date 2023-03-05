//go:build ignore_collisions
// +build ignore_collisions

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
这是一个基础控制语句的demo
*/

func main() {
	//判断语句if
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(21) - 10 //产生一个-10到10的随机数

	if x > 0 {
		fmt.Println("x大于0")
	} else if x < 0 {
		fmt.Println("x<0")
	} else {
		fmt.Println("x==0")
	}
	//switch
	switch x {
	case 1, 2, 3, 4, 5, 6:
		fmt.Println(x, "1")
	case 7, 8, 9, 10:
		fmt.Println(x, "2")
	case 0:
		fmt.Println(x, "3")
	default:
		//case都不符合的时候默认执行，在哪里并不影响
		fmt.Println("默认执行")
	case -10, -9, -8, -7:
		fmt.Println(x, "负数")

		/**
		在 Go 语言中，fallthrough 是一个关键字，它用于在 switch 语句中，将控制流穿过当前 case 语句块的结尾，继续执行下一个 case 语句块，而不论下一个 case 的条件是否匹配。

		fallthrough 的主要作用是让当前 case 的执行结果可以传递到下一个 case 中。当使用 fallthrough 时，程序将继续执行下一个 case 语句块的代码，而不会执行下一个 case 语句块的条件判断语句。

		需要注意的是，fallthrough 语句只能出现在 case 语句块的结尾，且只能在有下一个 case 语句块的情况下使用。如果当前 case 是最后一个 case，那么使用 fallthrough 将会导致编译错误。

		下面是一个示例，展示了如何使用 fallthrough 关键字：
		*/
		num := 2
		switch num {
		case 1:
			fmt.Println("one")
			fallthrough
		case 2:
			fmt.Println("two")
			fallthrough
		case 3:
			fmt.Println("three")
		default:
			fmt.Println("unknown")
		}
	}

}
