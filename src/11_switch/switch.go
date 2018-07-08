// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
switch
==================================================================
switch语句 相当于多个 if-else 判断语句的变体
------------------------------------------------------------------
语法：
	switch 初始语句; 判断参数 {
		case 判断值:
			// do something
		case 判断值:
			// do something
		case 判断值:
			// do something
		default:
			// do something
	}
------------------------------------------------------------------
switch语句 的初始语句和if的初始语句、for前置语句一致
switch语句 判断参数可以是任何可以比较的数据类型， case 无需为常量，且取值不必为整数
switch语句 default 相当于 else 结构体，当case 都不满足情况下执行
switch语句 只运行选定的 case，而非之后所有的 case，Go 在每个case后都自动添加break
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"runtime"
)

func os() {

	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			// Windows, openbsd
			fmt.Println("OS: ", os)
	}
}

func main() {

	os()
}
