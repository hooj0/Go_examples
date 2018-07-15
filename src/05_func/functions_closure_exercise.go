// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
func closure exercise
==================================================================
func闭包练习 斐波纳契闭包
------------------------------------------------------------------
利用闭包完成 斐波纳契闭包 的功能
------------------------------------------------------------------
------------------------------------------------------------------*/
package main

import "fmt"

func fibonacci() func() int {

	prev := 0
	current := 1
	ret := 0

	return func() int {

		ret = prev + current
		prev = current
		current = ret

		return ret
	}
}

func fibonacci2() func() int {

	arrays := []int { 0, 0, 1 }

	return func() int {

		arrays[2] = arrays[0] + arrays[1]

		arrays[0] = arrays[1]
		arrays[1] = arrays[2]

		return arrays[2]
	}
}

func main() {
	f := fibonacci2()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
