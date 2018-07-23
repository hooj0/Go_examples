// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-23
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
image
==================================================================
图像
------------------------------------------------------------------
接口定义：
	package image

	type Image interface {
		ColorModel() color.Model	//  颜色模式
		Bounds() Rectangle			// 	图片边界
		At(x, y int) color.Color	// 	某个点的颜色
	}
------------------------------------------------------------------
Bounds 方法的返回值 Rectangle 实际上是一个 image.Rectangle，它在 image 包中声明。
color.Color 和 color.Model 类型也是接口，
但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。
这些接口和类型由 image/color 包定义。
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"image"
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(img.Bounds())
	fmt.Println(img.At(0, 0).RGBA())
}
