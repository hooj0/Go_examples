// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-05
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
for switch break LABEL
==================================================================
for switch break 可以跳出循环或switch，并且可以调到指定的位置，但是不会再次执行循环
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"time"
)

func main() {


	fmt.Println("for start...")


	var size int8 = 0
	Loop:
		for n := 0; n < 5; n++ {

			sec := time.Now().Second()
			fmt.Println("sec: ", sec)
			switch {
				default:
					size = -1

				case sec < 30:
					size = 1

				case sec >= 30:
					size = 2
					break Loop
			}
		}

	fmt.Println("for end...")

	fmt.Println(size)

	// for start...
	// for end...
	// 2
}
