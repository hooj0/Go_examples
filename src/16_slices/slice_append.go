// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-10
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
slice append
==================================================================
append 追加元素，可以为切片添加新的元素
------------------------------------------------------------------
语法：
	func append(s []T, vs ...T) []T
	参数1：切片类型
	参数N：对应参数1切片类型的数据，随后追加进去的值
------------------------------------------------------------------
append 在原始数据基础追加新元素数据
append 追加新元素时，原始数组长度容量不够，会自动扩展。
append 扩展后的数组是一个全新的数组，相当于把原始数组和追加数组拷贝到新数组
------------------------------------------------------------------*/
package main

import "fmt"

func output(slices []int)  {
	fmt.Printf("len: %d, cap: %d, items: %v\n", len(slices), cap(slices), slices)
}

func main() {

	var arrays []int
	fmt.Printf("%T", arrays)
	output(arrays)		// len: 0, cap: 0, items: []

	// 向arrays追加2个新元素
	arrays = append(arrays, 1, 2)
	output(arrays)		// len: 2, cap: 2, items: [1 2]

	// 向arrays追加1个新元素
	arrays = append(arrays, 0)
	output(arrays)		// len: 3, cap: 4, items: [1 2 0]

	var numbers = []int{ 4, 5, 6, }
	// arrays = append(arrays, numbers)	// 编译不通过
	arrays = append(arrays, numbers...)	// 编译通过， len: 6, cap: 8, items: [1 2 0 4 5 6]
	output(arrays)
}
