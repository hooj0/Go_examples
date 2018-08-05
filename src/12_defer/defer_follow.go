// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-05
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
defer
==================================================================
defer 延迟执行
------------------------------------------------------------------
在存在defer行上，会先执行defer行中的参数求参，但 defer 执行时才会被求值
------------------------------------------------------------------*/
package main

import "fmt"

func trace(s string) string {
	fmt.Println("进入:", s)
	return s
}

func un(s string) {
	fmt.Println("离开:", s)
}

func a() {
	defer un(trace("a"))		// 优先执行trace函数进行求参，但推迟un函数的实际求值动作
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))		// 优先执行trace函数进行求参，但推迟un函数的实际求值动作
	fmt.Println("in b")
	a()
}

func main() {
	b()
}


// 进入: b
// in b
// 进入: a
// in a
// 离开: a
// 离开: b