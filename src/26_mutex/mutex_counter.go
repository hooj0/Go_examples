// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-30
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
mutex
==================================================================
互斥锁机制
------------------------------------------------------------------
互斥机制 是保证数据在同一时间不被多个线程访问共享，导致数据脏读/写的一种方法
------------------------------------------------------------------
用法：
Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法：
	Lock
		do something // 这中间的程序是安全的，不被共享
	Unlock

可以用 defer 语句来保证互斥锁一定会被解锁。
------------------------------------------------------------------
互斥机制 保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突

这里涉及的概念叫做 互斥（mutual exclusion） ，
我们通常使用 互斥锁（Mutex） 这一数据结构来提供这种机制。
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的
type SafeCounter struct {
	value 	map[string]int
	mux 	sync.Mutex
}

// Inc 增加给定 key 的计数器的值
func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()	// 从这里执行完后，只有一个 go routine 能访问s.value

	s.value[key]++

	s.mux.Unlock()
}

func (s *SafeCounter) Value(key string) int {

	s.mux.Lock()

	defer s.mux.Unlock()	// 不管结果如何，最终都执行解锁

	return s.value[key]		// 返回数据，同样每次只能一个线程访问 value
}

func main() {

	sc := SafeCounter{ value: make(map[string]int) }	// make int 默认为 0 值

	for i := 0; i < 1000; i++ {
		go sc.Inc("ok-key")				// 执行1000次，累加1000
	}

	time.Sleep(time.Second)

	fmt.Println(sc.Value("ok-key"))		// 1000

}
