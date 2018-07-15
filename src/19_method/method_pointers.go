// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
method pointer
==================================================================
指针接受者方法 方法接受者可以是一个指针对象，这样就为指针声明一个方法
------------------------------------------------------------------
指针接受者 对应之前接受者的 T 写法，指针接受者需要添加 *T
指针接受者 方法可以修改指针指向引用对象的值，这样修改接受者更方便

------------------------------------------------------------------*/
package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v Vertex) Sum() int {

	return v.x + v.y
}

// 当Add 接受者使用指针，会改变ref 对象的内存数据，改变Sum输出的结果
// 相反，如果不使用指针，这里的操作对原始Vertex值没有改变
func (v *Vertex) Add() {

	v.x += v.y
	v.y += v.x
}

func main() {

	ref := Vertex{ 2, 3 }

	// Add方法 使用指针
	fmt.Println("vertex: ", ref) 	// vertex:  {2 3}

	ref.Add()
	fmt.Println("vertex: ", ref) 	// vertex:  {5 8}

	fmt.Println("sum: ", ref.Sum())	// sum:  13

	// Add方法 不使用指针
	/*
	fmt.Println("vertex: ", ref) 	// vertex:  {2 3}

	ref.Add()
	fmt.Println("vertex: ", ref) 	// vertex:  {2 3}

	fmt.Println("sum: ", ref.Sum())	// sum:  5
	*/
}
