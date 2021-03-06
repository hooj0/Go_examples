// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------




/* ---------------------------------------------------------------
variable group
==================================================================
变量分组 主要用于同时声明多个不同类型变量
------------------------------------------------------------------
变量分组 同时声明多个类型变量，只需要一个var关键字
变量分组 在变量很多的情况下使用,可读性好、结构清晰。
变量分组 可以在func中或package中使用
变量分组 可以在分组中设置变量类型或初始值，但不能两者同时都不设置
------------------------------------------------------------------*/

package main

import "fmt"

// 声明一组变量
var (
	Data  string = "this is data str"
	Index int    = 12
	Flag  bool   = false

	Cource float64	// 可以没有初始值
)

// 声明一组变量，可以没有变量类型
var (
	Failure = true
	Max = 10
)

func main() {

	var (
		a int    = 1
		x int8   = 10
		z string = "hello"
	)

	fmt.Println("Data: ", Data, ", Index: ", Index, ", Flag: ", Flag)
	fmt.Println("a: ", a, ", x: ", x, ", z: ", z)

	fmt.Println("Failure: ", Failure, ", Max: ", Max, ", Cource: ", Cource)
}
