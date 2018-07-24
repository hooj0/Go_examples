// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-24
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
channel select
==================================================================
select 可以使 go runtime 同时等待多个通道操作
------------------------------------------------------------------
select 会阻塞某个分支流程可以继续执行为止，然后执行该分支
select 多个分支都准备好，都没有阻塞的情况下，会随机选择一个分支执行
------------------------------------------------------------------*/
package main

import "fmt"

func fibonacci(num, quit chan int)  {
	x, y := 0, 1

	for {
		select {
			case num <- x:
				x, y = y, x + y

			case <- quit:
				fmt.Println("quit...")
				// close(num)
				return
		}
	}
}

func main() {

	num := make(chan int)
	quit := make(chan int)

	// 匿名函数立即运行
	/*go func() {
		// 循环 10次，向 num 通道取数据；当循环完成后，将执行 quit
		for i := 0; i < 10; i++ {
			fmt.Println(<- num)
		}

		// 执行 quit ，向 quit 通道发送数据，select 接收到 quit 数据就执行退出循环
		quit <- 0
	}()*/

	go func() {
		for i := range num {
			fmt.Println(i)

			if i >= 100 {
				quit <- 0
			}
		}

	}()

	fibonacci(num, quit)
}
