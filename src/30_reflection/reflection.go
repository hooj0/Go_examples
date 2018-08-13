// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-05
// copyright by hoojo@2018
// --------------------------------------------------------------

/* ---------------------------------------------------------------
reflection 反射
==================================================================
反射 只是一种检查存储在接口变量中的“类型-值对”的机制
------------------------------------------------------------------
反射 reflect 包中的两种类型： Type 和 Value，这两种类型可用来访问接口变量的内容。

反射 还有两个简单的函数，叫做 reflect.TypeOf 和 reflect.ValueOf ，
		它们用于接口值中分别获取 reflect.Type 和 reflect.Value 。
		同样，从 reflect.Value 也能很容易地获取 reflect.Type
------------------------------------------------------------------
反射目的：

1、为了让 API 保持简单，Value 的“getter”和“setter”方法会在能够保存其值的最大类型上进行操作
2、反射对象的 Kind 描述了其基本类型，而给静态类型
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 2.34

	// 通过 TypeOf 获取数据类型
	fmt.Println("type:", reflect.TypeOf(x))			// type:  float64
	// 通过 ValueOf 获取数据类型，静态类型
	fmt.Println("value:", reflect.ValueOf(x).Type())	// value: float64

	v := reflect.ValueOf(x)
	// 静态类型
	fmt.Println("type:", v.Type())								// type: float64
	// 基本类型
	fmt.Println("Kind:", v.Kind())								// Kind: float64
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)	// kind is float64: true
	// 获取数据
	fmt.Println("value:", v.Float())								// value: 2.34


	var y uint8 = 'x'
	value := reflect.ValueOf(y)
	fmt.Println("type:", value.Type())                            // type: uint8
	fmt.Println("kind is uint8: ", value.Kind() == reflect.Uint8) // true.
	fmt.Println("uint:", value.Uint())							 // uint: 120
	y = uint8(value.Uint())                                       	 // v.Uint returns a uint64.
	fmt.Println("y:", y)											 // y: 120

	type MyInt int
	var a MyInt = 7
	v = reflect.ValueOf(a)
	fmt.Println("value:", v)				// value: 7
	// type 则可以描述 静态类型
	fmt.Println("type:", v.Type())		// type: main.MyInt
	// kind 永远描述基本类型（原始类型）
	fmt.Println("Kind:", v.Kind())		// Kind: int

}
