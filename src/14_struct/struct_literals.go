// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-09
// copyright by hoojo@2018
// --------------------------------------------------------------




/* ---------------------------------------------------------------
struct literals
==================================================================
结构体文法 通过直接列出字段的值来新分配一个结构体
------------------------------------------------------------------
结构体文法 在实例化结构体的时候可以直接通过结构体属性名称进行赋值
结构体文法 使用 Name: 语法可以仅列出部分字段进行设置值（字段名的顺序无关）
结构体文法 可以不设置任何属性值，构建一个仅仅只有默认值的结构体
结构体文法 在使用属性赋值的情况下，可以忽略顺序
结构体文法 特殊的前缀 & 返回一个指向结构体的指针
------------------------------------------------------------------*/
package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {

	// 只有默认值的结构体
	var v1 = Vertex{}
	fmt.Println("v1: ", v1)	// v1:  {0 0}

	// 按照属性顺序进行设置值
	var v2 = Vertex{2, 7}
	fmt.Println("v2: ", v2) 	// v2:  {2 7}

	// 为指定属性设置值
	var v3 = Vertex{Y: 5}
	fmt.Println("v3: ", v3) 	// v3:  {0 5}

	// 构建指针结构体
	var p = &Vertex{2, 7}
	fmt.Println("p: ", p) 	// p:  &{2 7}

	// 即便没有初始化，也是有零值
	var empty Vertex
	fmt.Printf("(%v, %T)\n", empty, empty)		// ({0 0}, main.Vertex)
	fmt.Println("X: ", empty.X, ", Y: ", empty.Y) 	// X:  0 , Y:  0
}
