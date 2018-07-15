// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
func value
==================================================================
func 也是一个数据类型，也是可以作为参数和返回值进行传递
------------------------------------------------------------------
func 作为参数传递到方法中，可以执行函数参数
func 作为返回值返回，能在其他地方执行该函数的返回值
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"strconv"
)

// 声明一个方法，参数为一个func，返回值是一个字符串
func output(fn func(int, string) string) string {

	// 执行参数传过来的func
	return fn(23, "Jack")
}


func main() {

	// 创建一个func，返回字符串，接收两个不同类型参数
	yourInfo := func(age int, name string) string {

		return "age: " + strconv.Itoa(age) + ", name: " + name
	}

	fmt.Println("jason info: ", yourInfo(18, "jason"))
	// jason info:  age: 18, name: jason

	// 传递方法直接传入方法名称即可
	fmt.Println("output: ", output(yourInfo))
	// output:  age: 23, name: Jack
}
