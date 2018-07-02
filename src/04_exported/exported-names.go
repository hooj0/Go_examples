// --------------------------------------------------------------
// author: hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create date: 2018-06-28
// copyright by hoojo @ 2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
exported
==================================================================
导出 主要用于是否对外公开，是否可以对其调用或访问
------------------------------------------------------------------
------------------------------------------------------------------
导出 在其他地方引入一个包后，访问包中的对象，会存在能否访问的权限问题。
	如果包中的大写字母开头的方法和属性，表示对外公开，可以进行访问或调用。
	如果包中的小写字母开头的方法和属性，表示对外不公开，不可以进行访问或调用。
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("math sqrt: ", math.Sqrt(2))
	fmt.Println("math PI: ", math.Pi)

	// undefined: fmt.parseArgNumber
	// fmt.Println(fmt.parseArgNumber("123")) // cannot refer to unexported name fmt.parseArgNumber
}
