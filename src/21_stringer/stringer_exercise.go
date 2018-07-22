// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-22
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
stringer exercise
==================================================================
stringer 练习 输出ip格式字符串
------------------------------------------------------------------*/
package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {

	return fmt.Sprintf("%v.%v.%v.%v \t", ip[0], ip[1], ip[2], ip[3])
}

func main() {

	local := IPAddr{ 127, 0, 0, 1 }
	remote := IPAddr{ 231, 23, 135, 61 }

	fmt.Println(local, remote)	// 127.0.0.1 	 231.23.135.61
}
