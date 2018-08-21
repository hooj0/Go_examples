// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-13
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
reflection
==================================================================
反射 修改源数据的值
------------------------------------------------------------------
	reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"reflect"
)

func set(x float64) {
	x = 2.22
}

func update()  {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.


	fmt.Println("type of p:", p.Type())			// type of p: *float64
	// 反射对象 p 并不是可设置的，不过我们也不想设置 p ， 而（实际上）是 *p
	fmt.Println("settability of p:", p.CanSet())	// settability of p: false

	// 调用 Value 的 Elem 方法，它会间接通过指针，并将结构保存到叫做 v 的反射值 Value 中：
	v := p.Elem()
	// 现在 v 是可设置的反射对象
	fmt.Println("settability of v:", v.CanSet())	// settability of v: true

	// 由于它代表 x ，因此最终我们可使用 v.SetFloat 来修改 x 的值：
	v.SetFloat(7.1)
	fmt.Println(v.Interface())	// 7.1
	fmt.Println(x)				// 7.1
}

func main() {

	var x float64 = 33.2

	// 将 x 的一份副本传入了 reflect.ValueOf
	// 因此该接口值也就作为传递给 reflect.ValueOf 的实参创建了一份 x 的副本，而非 x 本身
	v := reflect.ValueOf(x)

	fmt.Println("settability of v:", v.CanSet())	// settability of v: false

	if v.CanSet() {
		// 设置性是反射值 Value 的一种属性，而且并不是所有的反射值都拥有它
		// v 不可设置，假如能够成功执行，它也无法更新 x
		// 就算它能够更新存储在该反射值中的 x 的副本， x 本身也不会受影响
		v.SetFloat(7.1) // panic: reflect.Value.SetFloat using unaddressable value
	}

	// 将 x 传递给一个函数
	set(x)
	// set 不能修改 x ，因为我们传入的是值 x 的副本，而非 x 本身
	fmt.Println("x: ", x)	// x:  33.2

	// 想让 set 直接修改 x ，就必须向该函数传入 x 的地址（即指向 x 的指针）：
	// set(&x)

	update()
}
