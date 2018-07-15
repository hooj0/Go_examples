// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
func pointer
==================================================================
指针函数 在函数上使用指针，同之前method_pointers.go中一样的操作
------------------------------------------------------------------
指针函数 在传递参数或取值的地方，需要使用取地址符 & 和 取值 *
------------------------------------------------------------------*/
package main

import "fmt"

type Vertex struct {
	x, y int
}

func Sum(v Vertex) int {

	return v.x + v.y
}

func Add(v *Vertex) {

	v.x += v.y
	v.y += v.x
}

func main() {

	// 文法1：
	/*ref := Vertex{ 2, 3 }

	// Add方法 使用指针
	fmt.Println("vertex: ", ref) 	// vertex:  {2 3}

	Add(&ref)
	fmt.Println("vertex: ", ref) 	// vertex:  {5 8}

	fmt.Println("sum: ", Sum(ref))	// sum:  13
	*/

	// 文法2：
	ref := &Vertex{ 2, 3 }

	// Add方法 使用指针
	fmt.Println("vertex: ", *ref) 	// vertex:  {2 3}

	Add(ref)
	fmt.Println("vertex: ", *ref) 	// vertex:  {5 8}

	fmt.Println("sum: ", Sum(*ref))	// sum:  13


	// Add方法 不使用指针
	/*
	fmt.Println("vertex: ", ref) 	// vertex:  {2 3}

	Add(ref)	// 这里写法略有不同，去掉指向内存地址符号 &
	fmt.Println("vertex: ", ref) 	// vertex:  {2 3}

	fmt.Println("sum: ", Sum(ref))	// sum:  5
	*/

}
