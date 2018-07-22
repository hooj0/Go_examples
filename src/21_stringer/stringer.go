// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-22
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
stringer
==================================================================
stringer 是一个接口
stringer 一个可以用字符串描述自己的类型，有很多地方都用 stringer 接口来打印值
------------------------------------------------------------------
接口代码：
type Stringer interface {
    String() string
}
------------------------------------------------------------------*/
package main

import "fmt"

type User struct {
	Name string
	Age int
}

// 实现 Stringer 接口的 String() 方法
func (u User) String() string {

	return fmt.Sprintf("name: %v , age: %v ;", u.Name, u.Age)
}

func main() {

	foo := User{ "hoojo", 23 }
	bar := User{ "Jimmy", 52 }

	// 按照User对象stringer接口实现的方法String去输出数据
	fmt.Println(foo, bar)	// name: hoojo , age: 23 ; name: Jimmy , age: 52 ;

}
