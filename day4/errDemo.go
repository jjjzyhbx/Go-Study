package main

import (
	"errors"
	"fmt"
)

/**
type error interface{
	Error() string
}
该接口返回任何错误信息
*/

func returnError() (int, error) {
	return 1, errors.New("我返回了一个错误")
}

/*
*
自定义的错误类
*/
type myError struct {
	Massage string
}

/*
*
实现接口
*/
func (my myError) Error() string {
	return "我自定义的错误返回:" + my.Massage
}

/*
*
返回错误
*/
func aerror() error {
	return myError{"错误提示"}
}
func main() {
	_, err := returnError()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(aerror())
}
