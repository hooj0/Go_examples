// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
struct point
==================================================================
结构体指针 将结构体和指针结合使用，结构体字段可以通过结构体指针来访问
------------------------------------------------------------------
结构体指针 指针指向结构体，用指针操作结构体犹如利用结构体引用操作其本身
------------------------------------------------------------------*/
package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {

	var vert = Vertex{ 3, 7 }
	fmt.Println("vert: ", vert)					// vert:  {3 7}
	fmt.Println("x: ", vert.x, ", y: ", vert.y)	// x:  3 , y:  7

	// 指针指向 vert
	p := &vert
	fmt.Println("p: ", p)					// p:  &{3 7}
	fmt.Println("p: ", *p)					// p:  {3 7}
	fmt.Println("x: ", p.x, ", y: ", p.y)	// x:  3 , y:  7

	p.x = 666
	// 操作指针，进行设置值 和 读取值
	fmt.Println("p: ", p)		// p:  &{666 7}
	fmt.Println("p: ", *p)		// p:  {666 7}
	fmt.Println("vert: ", vert)	// vert:  {666 7}
}
