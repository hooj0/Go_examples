// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-22
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
error
==================================================================
错误 go语言中采用error表达错误的状态，error是一个接口
------------------------------------------------------------------
接口实现：
type error interface {
    Error() string
}
------------------------------------------------------------------
用法：
i, err := strconv.Atoi("42")	// 数据类型转换
if err != nil {		// 如果error 不为nil值，即转换错误
	// 进行错误处理
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}

// 否则 正常业务进行
fmt.Println("Converted integer:", i)

error 为 nil 时表示成功；非 nil 的 error 表示失败
------------------------------------------------------------------
错误 通常调用函数都会返回一个错误状态，通过利用 nil 来判断错误状态，来进行业务处理
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"time"
)

type Something struct {
	When time.Time
	What string
}

// 实现 error 接口的 Error 方法
func (s *Something) Error() string {
	return fmt.Sprintf("%v at %v o'clock", s.What, s.When)
}

func exec() error {
	return &Something{ time.Now(), "coding" }
}

func main() {

	foo := &Something{ time.Now(), "sleep" }

	fmt.Println("foo: ", foo, ", is nil: ", foo == nil)	// foo:  sleep at 2018-07-22 23:03:52.8286726 +0800 CST m=+0.000000001 o'clock , is nil:  false

	if err := exec(); err != nil {
		fmt.Println(err)	// coding at 2018-07-22 23:02:20.9687924 +0800 CST m=+0.024600501 o'clock
	}
}
