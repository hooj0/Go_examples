// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
basic types
==================================================================
基本数据类型 go lang 中的数据类型
------------------------------------------------------------------
基本数据类型 在语言中数据声明时所用的类型
基本数据类型 不同的类型用来存储不同返回的数据
基本数据类型 int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
		   当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
------------------------------------------------------------------
// 布尔类型，true / false
bool

// 字符串类型
string

// 数字类型，整数、长整数
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

// 字节类型
byte // uint8 的别名

rune // int32 的别名
     // 表示一个 Unicode 码点

// 浮点型，小数点数字
float32 float64

// 复杂类型
complex64 complex128
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Successed bool       = false
	MaxInt    uint64     = 1<<64 - 1
	z         complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("Successed: ", Successed, ", MaxInt: ", MaxInt, ", z: ", z)
}
