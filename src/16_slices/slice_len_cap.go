// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-10
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
slice length cap
==================================================================
切片 有长度和容量值限制
------------------------------------------------------------------
切片 长度就是切片的元素数量
切片 容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
切片 长度的获取方式：len(slices)
切片 容量的获取方式：cap(slices)
切片 可以重新切片来扩展切片容量大小
------------------------------------------------------------------*/

package main

import "fmt"

func output(slices []int)  {
	fmt.Printf("len: %d, cap: %d, items: %v\n", len(slices), cap(slices), slices)
}

func main() {

	arrays := []int { 1, 2, 3, 4, 5, 6, 7 }

	output(arrays)		// len: 7, cap: 7, items: [1 2 3 4 5 6 7]

	// 长度为0， 但是容量却是数组的length
	output(arrays[:0])	// len: 0, cap: 7, items: []
	// 切片是4个元素
	output(arrays[:4])	// len: 4, cap: 7, items: [1 2 3 4]
	// 去掉前面2个元素
	output(arrays[2:])	// len: 5, cap: 5, items: [3 4 5 6 7]
}
