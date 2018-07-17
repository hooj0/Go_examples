// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-17
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
interface value nil
==================================================================
接口值nil 没有具体实例化的实现对象的引用值
------------------------------------------------------------------
接口值nil 即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用
		 在接口实现没有初始化实例化的情况下，接口值的方法也能被调用
接口值nil 没有实例化的类型调用接口方法，Go也能优雅的处理
------------------------------------------------------------------
------------------------------------------------------------------*/
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("t is nil")
		return
	}
	fmt.Println("s: ", t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {

	var i I
	var t *T	// 取值，指针类型; 由于指针没有指向的内存地址，所有为nil
	var t2 T	// 零值，struct类型; 相当于 t2 = { S: "" } 一个零值的struct

	fmt.Printf("(%v, %T)\n", t, t)		// (<nil>, *main.T)
	fmt.Printf("(%v, %T)\n", t2, t2)		// ({}, main.T)
	fmt.Println("------------------------------------")

	i = t
	describe(i)		// (<nil>, *main.T)
	i.M()			// t is nil

	i = &T{ "hello" }
	describe(i)		// (&{hello}, *main.T)
	i.M()			// s:  hello

	// -----------------------------------------------
	// 不一样的结果
	/*
	var i I
	var t T

	i = &t
	describe(i)		// (&{}, *main.T)
	i.M()			// s:
	*/
}
