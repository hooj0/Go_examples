// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-16
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
interface
==================================================================
接口 由一组方法签名组成的集合体
------------------------------------------------------------------
语法：
	type 接口名称 interface {
		各种方法(x int)
		更多方法() int
	}
------------------------------------------------------------------
接口 类型的变量可以保存任何实现了接口方法的值
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"math"
)

// 定义一个接口：Abser
type Abser interface {

	// 接口中的方法
	Abs() float64
} 

// 定义一个 float 类型
type MyFloat float64

// 声明一个 接受者为 MyFloat 的方法
func (f MyFloat) Abs() float64 {

	if f > 0 {
		return  float64(-f)
	}

	return float64(f)
}

// 定义一个结构体
type Vertex struct {
	X, Y float64
}

// 声明一个接受指针结构体接受者的方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {


	// 声明 接口 变量
	var abser Abser

	// 实例化一个 MyFloat 对象
	f := MyFloat(-math.Sqrt2)

	// 实例化一个结构体对象
	v := Vertex{ 2, 4 }

	// 将 MyFloat 赋值给接口变量，就完成了接口实现
	abser = f	// MyFloat 实现了 Abser，MyFloat是Abs() 方法接收者

	fmt.Println(abser.Abs())	// -1.4142135623730951

	// 将 Vertex 引用 赋值给 接口变量，就完成了接口实现
	abser = &v	// *Vertex 实现了 Abser，*Vertex 是Abs() 方法接收者

	// v 是一个 Vertex（而不是 *Vertex） 所以没有实现 Abser。
	// abser = v // 编译不通过

	fmt.Println(abser.Abs())	// 4.47213595499958
}
