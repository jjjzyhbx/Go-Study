package main

import "fmt"

/**
这是一个接口的小demo
*/

/**
在Go语言中，接口是一组方法的集合，可以用来定义对象的行为和规范。
接口定义了对象应该具有哪些方法，但是并不实现这些方法，具体的实现由对象来完成。
接口可以实现多态性，也可以用于解耦合和模块化程序设计。
接口定义如下
type MyInterface interface {
    Method1(param1 type1, param2 type2) returnType1
    Method2(param3 type3) returnType2
    ...
}

*/

type runTimeThread interface {
	run()
}
type myClass struct {
	massage string
	status  string
}
type classLoader struct {
	loaderMassage string
	loader        interface{} //空接口，表示任意类型
}

func (my myClass) run() {
	fmt.Println("this is myClass", my)
}
func (loader classLoader) run() {
	fmt.Println("this is classLoader", loader)
}

/**
基于空接口的函数参数，可以传递任何类型
*/
func Prints(a interface{}) {
	fmt.Println(a)
}
func main() {
	/**
	基于接口实现的一个简单多态,
	实现接口里的全部方法，就实现了该接口
	*/
	my := myClass{massage: "myClassMassage", status: "ok"}
	loader := classLoader{"classLoaderMassage", "this loader"}
	var r runTimeThread = my
	r.run()
	r = loader
	r.run()

	//空接口
	var i interface{} //该接口可以表示任何类型
	fmt.Println(i)
	Prints("dshbci")
	Prints(1)
	Prints(12.34)
	Prints(true)

}
