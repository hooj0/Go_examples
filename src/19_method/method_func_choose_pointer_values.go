// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-16
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
func & method choose pointer | values
==================================================================
函数和方法选择指针和值作为接收者
------------------------------------------------------------------
1、方法选择指针接受者，可以轻松传入指针或非指针接收者参数，虽然函数也能适当改变满足
2、方法选择指针接受者，可以指向引用的值，方便在方法中进行修改
3、这样能避免复制修改的值，当接受者是赋值类型的情况下，会很高效，但函数却需要复制修改结果
4、方法的接受者可以是指针或普通非指针，但不建议混用
------------------------------------------------------------------*/
package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func (v *Vertex) Abs() int {
	return v.x * v.x + v.y * v.y
}

func main() {

	// 取址 和 普通 写法都可以正常运行，结果一致
	vert := &Vertex{ 3, 4 }
	// vert := Vertex{ 3, 4 }

	fmt.Printf("before scale: %+v, v.Abs: %v \n", vert, vert.Abs())	// before scale: &{x:3 y:4}, v.Abs: 25

	vert.Scale(3)
	fmt.Printf("after scale: %+v, v.Abs: %v \n", vert, vert.Abs())	// after scale: &{x:9 y:12}, v.Abs: 225
}
