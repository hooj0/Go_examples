// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-16
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
func pointer redirect
==================================================================
方法（函数）指针与重定向
	在使用指针接受者的情况下，需要使用指针地址符改变内存地址指向
------------------------------------------------------------------
函数指针重定向 在接收指针函数的情况下，都必须接收一个指针参数
方法指针重定向 在方法的参数为指针接受者的情况下，既可以是指针参数也可以是非指针
------------------------------------------------------------------
在底层Go语言会将指针接受者参数进行按地址转换，即使是非指针传参情况也能正常调用
------------------------------------------------------------------*/
package main

import "fmt"

type Vertex struct {
	X, Y int
}

// 方法
func (v *Vertex) ScaleMethod(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// 函数
func ScaleFunc(v *Vertex, i int)  {
	v.X = v.X * i
	v.Y = v.Y * i
}

func main() {

	// 文法1：
	// --------------------------------------------------------------------
	ref := Vertex{3, 6 }
	fmt.Println("vertex: ", ref)		// vertex:  {3 6}

	ref.ScaleMethod(5)
	fmt.Println("run method vertex: ", ref)	// run method vertex:  {15 30}

	// ScaleFunc(ref, 6) 	// 编译错误
	ScaleFunc(&ref, 6)  	// 编译成功
	fmt.Println("run func vertex: ", ref)	// run func vertex:  {90 180}

	// 文法2：
	// --------------------------------------------------------------------
	ref2 := &Vertex{3, 6 }
	fmt.Println("vertex: ", ref2)		// vertex:  &{3 6}

	// 由于 ScaleMethod 有一个指针接收者，为方便起见，Go 会将语句 v.ScaleMethod(5) 解释为 (&v).Scale(5)
	ref2.ScaleMethod(5)
	fmt.Println("run method vertex: ", ref2)	// run method vertex:  &{15 30}

	// ScaleFunc(ref, 6) 	// 编译错误
	ScaleFunc(ref2, 6)  	// 编译成功
	fmt.Println("run func vertex: ", ref2)	// run func vertex:  &{90 180}
}
