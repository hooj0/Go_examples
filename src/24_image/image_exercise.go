// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-23
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
image exercise
==================================================================
图像 练习
------------------------------------------------------------------
------------------------------------------------------------------
------------------------------------------------------------------*/
package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	width int
	height int
}


func (img Image) Bounds() image.Rectangle {

	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) ColorModel() color.Model {

	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {

	return color.RGBA{uint8(x),uint8(y),uint8(255),uint8(255) }
}

func main() {
	m := Image{50, 50 }
	pic.ShowImage(m)
}
