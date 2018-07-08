// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------




/* ---------------------------------------------------------------
data type conversion
==================================================================
类型转换 用于将一个类型转换至另一个类型
------------------------------------------------------------------
表达式 T(v) 将值 v 转换为类型 T
------------------------------------------------------------------
类型转换 不同的类型之间进行转换后，进行同类型的运算操作
类型转换 存在类型转换错误异常
类型转换 存在类型转换数据溢出
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"math"
)

func main() {

	// var 数值的转换：
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("i: %v, f: %v, u: %v \n", i, f, u)

	// 短变量 数值转换
	index := 23
	score := float64(index)
	money := uint(score)

	// 字符串不能转换为整型
	//item := "hi"
	//result := int8(item)

	fmt.Printf("index: %v, score: %v, money: %v \n", index, score, money)

	var x, y int = 3, 4
	var ret float64 = math.Sqrt(float64(x * x + y * y))
	var z uint = uint(ret)

	fmt.Println("x: ", x, ", y: ", y, ", ret:, ", ret, ", z: ", z)
}
