// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
switch executed order
==================================================================
switch语句 执行顺序自上而下，直到匹配到指定的case或default为止
------------------------------------------------------------------
switch语句 default不是必选的，当没有匹配到将不会有默认结构体执行
switch语句 case 语句中的条件语句在匹配时，每次会被执行。
switch语句 case 语句中匹配成功的会执行结构体，后面没有匹配的不会执行
------------------------------------------------------------------*/

package main

import "fmt"

func num() int {

	fmt.Println("exec num func")

	return 5
}

func like(x int) {

	fmt.Println("input: ", x)
	switch x {
		case 1:
			fmt.Println("like orange")
		case 2:
			fmt.Println("like apple")
		case num():
			fmt.Println("like Durian")
		case 5:
			fmt.Println("like mango")

	}
}

func main() {

	// 不会执行到 num
	like(1)

	// 会执行到 num
	like(3)

	// 会执行到 num，相当于 所有 case 判断语句都被执行一次
	like(10)
}
