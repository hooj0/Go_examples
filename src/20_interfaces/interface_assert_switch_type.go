// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-22
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
interface assert type switch
==================================================================
接口断言之类型选择 可以通过断言判断值的类型，来选择正确的业务流程
------------------------------------------------------------------
语法：
	switch v := i.(type) {
	case T:
	    // v 的类型为 T
	case S:
	    // v 的类型为 S
	default:
	    // 没有匹配，v 与 i 的类型相同
	}
------------------------------------------------------------------
此选择语句判断接口值 i 保存的值类型是 T 还是 S。
在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值。
在默认（即没有匹配）的情况下，变量 v 与 i 的接口类型和值相同。
------------------------------------------------------------------
类型选择 采用switch语法形式进行流程选择，在匹配的流程分支中进行业务处理
类型选择 中的类型还是采用 i.(T) 的语法形式，但 T 换成 type 类型
------------------------------------------------------------------*/
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
		case int:
			fmt.Printf("int param: %v, exp: %v\n", v, v*2)
		case string:
			fmt.Printf("string param: %v, length: %v \n", v, len(v))
		default:
			fmt.Printf("unknown type: %T", v)
	}
}

func main() {

	do(11)		// int param: 11, exp: 22

	do("jack")	// string param: jack, length: 4

	do(2.2)		// unknown type: float64
}
