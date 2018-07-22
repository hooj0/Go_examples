// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-22
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
interface assert
==================================================================
接口类型断言 用来判断接口的实现类型，提供了访问接口值底层具体值的方式
------------------------------------------------------------------
语法：
	var i interface = "a"
	t := i.(T) 	// 断言 接口 i 保存的具体类型为 T，并将类型 T 赋值给 变量 t
				// 若 i 接口并未保存的类型不是 T， 就会触发错误

	t, ok := i.(T)  // 可以接收两个参数，能够判断是否为类型 T
					// 若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。
					// 否则，ok 将为 false 而 t 将为 T 类型的零值，程序并不会产生恐慌。
------------------------------------------------------------------
请注意：这种语法和读取一个映射时的相同之处
------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	var i interface{} = "interface assert"
	s := i.(string)
	fmt.Println("interface type: ", s)	// interface type:  interface assert

	s, ok := i.(string)
	fmt.Println("interface type: ", s, ", ok: ", ok)	// interface type:  interface assert , ok:  true

	// s, ok = i.(float64)			// 编译不通过
	var s2, ok2 = i.(float64)		// 不是 float64 类型，s2 的值为零值
	fmt.Println("interface type: ", s2, ", ok: ", ok2)	// interface type:  0 , ok:  false

	s3 := i.(int)	// error： panic: interface conversion: interface {} is string, not int
	fmt.Println("interface type: ", s3)
}
