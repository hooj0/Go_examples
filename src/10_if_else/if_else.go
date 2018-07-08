// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
if else
==================================================================
判断语句 if 和 else 必须执行一个结构体
------------------------------------------------------------------
语法：
	if 判断语句 {
		// do something
	} else {
		// do something
	}
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"math"
)

func max(f float64, max float64) float64  {

	// 在条件语句之前运行一段短语句
	if x := math.Round(f); x < max {
		return x
	} else {
		fmt.Println("x >= max")
	}

	return max
}

func main() {

	fmt.Println("max: ", max(2.23, 10.0))
	fmt.Println("max: ", max(10.23, 10.0))
}
