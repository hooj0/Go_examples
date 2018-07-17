// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-17
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
interface values
==================================================================
接口值 在接口内部，接口值可以看作包含值和类型的元组
------------------------------------------------------------------
接口值 保存了一个具体底层类型的实际值
接口值 调用方法时，会执行底层类型的同名方法
接口值 就是一个接口对象实现的类型的具体实例引用
------------------------------------------------------------------*/
package main

import "fmt"

type IUser interface {
	MOutput()
}

type SUserInfo struct {
	name string
}

func (user *SUserInfo) MOutput() {
	fmt.Println("user name: ", user.name)
}

type FUserScore float64

func (f FUserScore) MOutput() {
	fmt.Println("user score: ", float64(f))
}

func describe(i IUser) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {

	// 这里的 user 就是接口值，每次赋值都是一个新的接口，指向底层结构类型的底层方法
	var user IUser = &SUserInfo{ "jack" }
	describe(user)		// (&{jack}, *main.SUserInfo)
	user.MOutput()		// user name:  jack

	user = FUserScore(66.3)
	describe(user)		// (66.3, main.FUserScore)
	user.MOutput()		// user score:  66.3
}
