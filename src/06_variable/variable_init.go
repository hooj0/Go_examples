// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-03
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
variable init
==================================================================
变量 初始化是在变量声明期间设定的值
------------------------------------------------------------------
变量 初始化期间可以同时为多个变量设置默认值，变量的数量和设置值的数量一致
变量 初始化可以在package和func中进行
变量 如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
------------------------------------------------------------------*/

package main

import (
	"flag"
	"fmt"
		"os"
)

// 在package中进行初始化
var foo, bar, data int8 = 1, 3, 54

// 可以设置表达式，在后期运算时求值
var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

// init 方法是初始化方法，可以初始化一些变量的值
// init 函数还常被用在程序真正开始执行前，检验或校正程序的状态。
func init() {
	if user == "" {
		// 输出错误日志，抛出错误
		// log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	// gopath 可通过命令行中的 --gopath 标记覆盖掉。
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

func main() {

	// 在 func 中进行声明初始化
	var x, y int = 2, 6

	// 声明变量，不再需要设置类型，可以直接赋值
	var a, i, result = true, false, false

	fmt.Println("x: ", x, ", y: ", y) // x:  2 , y:  6
	fmt.Println("a: ", a, ", i: ", i, ", result: ", result) // a:  true , i:  false , result:  false
	fmt.Println("foo: ", foo, ", bar: ", bar, ", data: ", data) // foo:  1 , bar:  3 , data:  54

	fmt.Println("home:", home, ", user:", user, ", gopath:", gopath)
}
