// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------




/* ---------------------------------------------------------------
const
==================================================================
常量 用于声明一个固定不变，全局可用的变量数据
------------------------------------------------------------------
const PI = 3.14159267
------------------------------------------------------------------
常量 声明与变量类似，只不过是使用 const 关键字
常量 可用被外部访问
常量 可以是字符、字符串、布尔值或数值类型
常量 不能用 := 语法声明。
------------------------------------------------------------------*/


package main

import "fmt"

// 常量声明
const Pi  = 3.14159267
const MAX_VALUE  = 10122111.2

// 常量组
const (
	a = 1
	b = true
	c = "str"
)


type ByteSize float64

// 枚举常量使用枚举器 iota 创建
// 由于 iota 可为表达式的一部分，而表达式可以被隐式地重复，这样也就更容易构建复杂的值的集合
const (
    _           = iota // 通过赋予空白标识符来忽略第一个值
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

type Int int

const (
	_ 		= iota		// 0， 	iota 默认从0开始递增
	A Int 	= 1 + iota	// 1 + 1
	B					// 1 + 2
	C					// 1 + 3
	D
)

func (b ByteSize) String() string {
    switch {
		case b >= YB:
			return fmt.Sprintf("%.2fYB", b/YB)
		case b >= ZB:
			return fmt.Sprintf("%.2fZB", b/ZB)
		case b >= EB:
			return fmt.Sprintf("%.2fEB", b/EB)
		case b >= PB:
			return fmt.Sprintf("%.2fPB", b/PB)
		case b >= TB:
			return fmt.Sprintf("%.2fTB", b/TB)
		case b >= GB:
			return fmt.Sprintf("%.2fGB", b/GB)
		case b >= MB:
			return fmt.Sprintf("%.2fMB", b/MB)
		case b >= KB:
			return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}

func main() {

	fmt.Printf("Pi: %v, MAX_VALUE: %v \n", Pi, MAX_VALUE)
	fmt.Printf("a: %v, b: %v, c: %v \n", a, b, c)

	// 常量声明
	const Valid  = false
	fmt.Println("Valid: ", Valid)

	fmt.Println("GB:", 1.22 * GB)	// GB: 1.22GB

	fmt.Println("A:", A)				// A: 2
	fmt.Println("D:", D)				// D: 5
}
