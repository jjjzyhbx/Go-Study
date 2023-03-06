package main

import (
	"fmt"
	"os"
)

/**
这是一个基础的go文件操作
*/

func main() {
	//文件的创建与打开
	file, err := os.Create("test.txt") //在当前项目目录下创建一个test

	if err != nil {
		fmt.Println("打开失败")
		fmt.Println(err)
		return
	} else {
		fmt.Println("打开成功")
		fmt.Println(file)
	}
	defer file.Close()
	//文件的打开
	fil, err := os.Open("tes.txt") //在当前项目目录下创建一个test

	if err != nil {
		fmt.Println("打开失败")
		fmt.Println(err)
		return
	} else {
		fmt.Println("打开成功")
		fmt.Println(fil)
	}

	//文件的读取与写入
	//Read
	buf := make([]byte, 1024)
	n, er := fil.Read(buf)
	if er != nil {
		fmt.Println(er, "读取失败")
	} else {
		fmt.Println(string(buf[:n]))
	}
	//write
	data := []byte("坤坤打篮球")
	_, e := file.Write(data)
	if e != nil {
		fmt.Println(e, "写入失败")
	} else {
		fmt.Println("ok")
	}

	//文件的关闭
	defer fil.Close()
	defer file.Close()

}
