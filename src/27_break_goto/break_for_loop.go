// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-05
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
for break LABEL
==================================================================
跳出循环，并跳出到指定的label位置，跳出后在LABEL位置继续执行，并不再进入当前for循环
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"time"
)

func simpleFor()  {
	fmt.Println("simple for loop start...")

	LOOP:
		for i := 0; i < 5; i++ {

			fmt.Println("i -> ", i)

			sec := time.Now().Second()
			if sec > 30 {
				break LOOP
			}
		}

	fmt.Println("simple for loop end...")
}

func complexFor()  {

	fmt.Println("complex for loop start...")

	LOOP:
		for n := 0; n < 5; n++ {
			fmt.Println("n: ", n)

			for i := 0; i < 5; i++ {

				fmt.Println("i -> ", i)

				sec := time.Now().Second()
				if sec > 30 {
					break LOOP
				}
			}
		}

	fmt.Println("complex for loop end...")
}
func main() {

	simpleFor()
	complexFor()
}
