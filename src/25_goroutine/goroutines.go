// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-24
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
go routine
==================================================================
go routine 是由 Go 运行时管理的轻量级线程。
------------------------------------------------------------------
语法：
	go f(x, y, z)		// 会启动一个新的 Go 程并执行
f (x, y, z) 的求值发生在当前的Go线程中，而 f 执行则会在新的线程中执行
------------------------------------------------------------------
go routine 在相同的地址空间中运行，因此访问共享内存的数据时必须同步。
go routine 同步可以利用 sync 关键字，它提供了共享数据同步的能力
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"time"
)

func say(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)	// 休眠 100 毫秒
		fmt.Println("str: ", str)
	}
}

func main() {

	go say("hallo go routine")	// 不能在当前程序退出的情况下执行go routine
	say("hello!")	// 当前程序执行完成，go线程也立即完成
}

// str:  hallo go routine
// str:  hello!
// str:  hello!
// str:  hallo go routine
// str:  hallo go routine
// str:  hello!
// str:  hallo go routine
// str:  hello!
// str:  hallo go routine
// str:  hello!