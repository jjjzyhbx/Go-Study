/*
*
Go语言中的并发是通过goroutine和channel来实现的。
goroutine是一种轻量级的线程，它由Go语言的运行时(runtime)来调度和管理。
channel是一种用于goroutine之间通信和同步的机制。
定义格式为 var ch chan T ，T为任意类型
*/
package main

import "fmt"

func main() {
	//用例1
	ch := make(chan string, 1) //创建了一个字符串通道 ch
	// 创建一个带缓冲的字符串通道，缓冲区大小为1
	//ch <- x  向channel中发送数据x
	//x := <-ch 从channel中接收数据，将其赋值给变量x
	go func(str string) {
		ch <- str
	}("xiaoming")
	fmt.Println(<-ch)

	//缓冲通道
	/**
	缓冲通道是一种具有一定容量的channel，可以在不阻塞发送操作的情况下存储多个元素。
	可以使用内置的make函数创建一个缓冲通道，并通过指定缓冲区的容量来指定缓冲区的大小，例如：
	*/
	c2 := make(chan int, 10) // 创建一个容量为10的缓冲通道
	fmt.Println(c2)

	//单向通道

	/**
	单向通道是指只允许发送或接收操作的channel。
	可以使用类型转换将一个双向通道转换为单向通道
	*/
	sendCh := ch   // 转换为发送操作通道
	recvCh := <-ch // 转换为接收操作通道
	fmt.Println(sendCh, recvCh)

	//select 语句，

	/**
	3. select语句
	select语句可以用于同时监听多个channel上的事件，并在其中任意一个channel上发生事件时执行对应的操作。
	select语句类似于switch语句，但是它的每个case子句都是一个channel操作，
	这是一个适用于通道的Switch语句
	*/
	select {
	case x := <-ch:
		fmt.Println("received from ch:", x)
	case y := <-ch:
		fmt.Println("received from ch:", y)
	default:
		fmt.Println("no activity")
	}

}
