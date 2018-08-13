// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-13
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
reflection interface
==================================================================
接口反射	从反射对象可反射出接口值
------------------------------------------------------------------
reflect.Value 使用 Interface 方法还原其接口值，该方法会将类型与值的信息打包成接口表示，并返回其结果
------------------------------------------------------------------
接口：
	func (v Value) Interface() interface {
	}
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 2.34
	// 获取 Value 对象
	v := reflect.ValueOf(x)

	// Interface 将 v 的值返回成 interface{}
	// func (v Value) Interface() interface {}

	// y 将具有 float64 类型
	var y = v.Interface().(float64)

	// 反射对象 v 所表示的 float64 值
	fmt.Printf("y: %v, v: %T \n", y, y)				// y: 2.34, v: float64
	// Println 可以接收 interface的参数
	fmt.Println(v.Interface())								// 2.34
	fmt.Printf("value is %7.1e\n", v.Interface())	// value is 2.3e+00

}
