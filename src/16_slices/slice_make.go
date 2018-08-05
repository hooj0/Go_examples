// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-10
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
make slice
==================================================================
make 使用内建函数 make 创建切片，一种创建动态数组的方式。
------------------------------------------------------------------
语法：
	a := make([]int, 5)  	// len(a)=5
	a := make([]int, 0, 5)  // len(b)=0, cap(b)=5
------------------------------------------------------------------
make 创建切片会首先创建一个数组，数组中存储都是nil值，同时返回数组的切片
make 创建切片，可以设定长度和容量的值
make 还能“只”用于创建切片(slice)、映射(map)和信道(chan)
------------------------------------------------------------------*/

package main

import "fmt"

func output(typeName string, slices []int)  {
	fmt.Printf("%s -> len: %d, cap: %d, items: %v\n", typeName, len(slices), cap(slices), slices)
}

func main() {

	// 构建 长度 5，容量 0，值为 0 的切片
	a := make([]int, 5)
	output("a", a)	// a -> len: 5, cap: 5, items: [0 0 0 0 0]

	// 构建 长度 0，容量 5，值为 nil 的切片
	b := make([]int, 0, 5)
	output("b", b)	// b -> len: 0, cap: 5, items: []

	// 切片扩充长度 2，值 从 nil 变为 0
	c := b[:2]
	output("c", c)	// c -> len: 2, cap: 5, items: [0 0]

	// 切片扩充长度 3，值 从 nil 变为 0
	d := c[2:5]
	output("d", d)	// d -> len: 3, cap: 3, items: [0 0 0]
}
