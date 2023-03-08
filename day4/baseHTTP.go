package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
这也是一个http基础知识
*/

/*
*
下面是一个实现Handler的接口
*/
type Engine struct {
	a interface{}
}

func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
		engine.a = req.URL
		fmt.Println(engine.a)
	case "/hello":
		for k, v := range req.Header {
			fmt.Println(k, v)
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)

		}
		engine.a = req.URL
		fmt.Println(engine.a)
	default:
		fmt.Println("404")
		fmt.Fprintf(w, "404 %s\n", req.URL)
		engine.a = req.URL
		fmt.Println(engine.a)
	}
}
func main() {
	//通过Engine 实例可以操作
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":8080", engine))
}
