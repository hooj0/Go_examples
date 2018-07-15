// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
slice 练习
==================================================================
根据输入的切片长度、切片元素长度参数，构建一个颜色取值在灰色或蓝色渐变的色块图片
------------------------------------------------------------------
将image后面的数据输入到浏览器地址栏上预览
http://data:image/gif;base64, 填base64字符串
------------------------------------------------------------------*/

package main

import (
	"golang.org/x/tour/pic"
)

func Pic(slice_size, item_size int) [][]uint8 {
	pic_slices := make([][]uint8, slice_size)

	for x := range pic_slices {
		pic_slices[x] = make([]uint8, item_size)

		for y := range pic_slices[x] {
			// 渐变效果
			pic_slices[x][y] = uint8((x + y) / 2)
			// 4 个大块对称效果
			// pic_slices[x][y] = uint8(x * y)
			// pic_slices[x][y] = uint8(x ^ y)
			// 非对称
			// pic_slices[x][y] = uint8(x % (y + 1))
		}
	}

	return pic_slices
}

func main() {

	pic.Show(Pic)
}
