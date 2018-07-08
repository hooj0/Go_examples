// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
variable default
==================================================================
变量默认值 在声明变量时没有明确指定初始值，系统会默认给予初始默认值
------------------------------------------------------------------
变量默认值 声明变量只提供类型，不提供值的情况
------------------------------------------------------------------*/

package main

import "fmt"


// 声明变量，没有设置初始值
var Data string
var Range int

func main() {

	// 声明变量，没有设置初始值
	var i 		int
	var index 	uint64
	var proto 	byte
	var str 	string
	var score	float64
	var flag 	bool

	// Printf 可以对字符串进行格式化输出
	fmt.Printf("Data: %q, Range: %v \n", Data, Range)

	fmt.Printf("i: %v, index: %v, proto: %v, str: %q, score: %v, flag: %v \n", i, index, proto, str, score, flag)
}
