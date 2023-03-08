package main

import (
	"fmt"
	"net/http"
)

/*
*
这是一个http标准库的一个学习，
*/
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "你打开了一个网络web服务")
}
func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Println("我访问了该路由")
}

func main() {

	/**
	main()函数通过代码http.ListenAndServe(":8080“,nil)启动一个8080端口的服务器。
	此时在网页中输入http://localhost:8080/hello，就可以启动
	*/
	http.HandleFunc("/hello", hello)
	//设置hello 路由，绑定hello方法
	http.HandleFunc("/", indexFunc)
	//设置空路由，绑定indexFunc 函数

	//启动监听 8080,
	http.ListenAndServe(":8080", nil)
	// 在浏览器中输入 http://localhost:8080/hello，
	//就可以访问该页面

	/**
	ListenAndServe()函数有两个参数，当前监听的端口号和事件处理器Handler
	事件处理器接口如下
	type Handler interface {
	     ServeHTTP(ResponseWriter, *Request)
	 }
	http包中实现该方法如下
	type HandlerFunc func(ResponseWriter,*Request)
	func (f HandlerFunc ) ServeHTTP (w ResponseWrite,r *Request){
	f(w,r)
	}
	*/

}
