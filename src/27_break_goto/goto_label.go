// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-05
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
goto
==================================================================
跳转语句 goto语句可以跳转到本函数内的某个标签
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"time"
)

func main() {

	sec := time.Now().Second()
	fmt.Println("sec: ", sec)

	label1:
		fmt.Println("go to label 1")
		if sec > 10 && sec < 30 {
			goto label2
		} else {
			goto label4
		}
	label2:
		fmt.Println("go to label 2")
		if sec > 30 && sec < 50 {
			goto label3
		} else {
			goto label4
		}
	label3:
		fmt.Println("go to label 3")
		if sec < 10 {
			goto label1
		} else {
			goto label4
		}
	label4:
		fmt.Println("go to label 4")

}
