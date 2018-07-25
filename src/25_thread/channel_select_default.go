// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-24
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
select default
==================================================================
select默认值 同switch一样，select也有默认值，当select的通道都没有准备好，default就执行
------------------------------------------------------------------
select默认值 如果期望在发送或接收时，不发生任何阻塞，可以使用default分支
------------------------------------------------------------------
语法：
	select {
		case i := <-c:
			// 使用 i
		default:
			// 从 c 中接收会阻塞时执行
	}
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)	// 100 毫秒触发一次
	boom := time.After(500 * time.Millisecond)	// 500 毫秒后触发

	for {
		select {
			case <- tick:
				fmt.Println("tick")
			case <- boom:
				fmt.Println("boom")
			default:
				fmt.Println("   .")
				time.Sleep(50 * time.Millisecond)	// 休眠 50 毫秒
		}
	}
}
