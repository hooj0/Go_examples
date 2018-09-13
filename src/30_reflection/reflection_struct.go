// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-09-13
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
reflection struct
==================================================================
反射结构体 通过结构体的地址， 修改它的字段
------------------------------------------------------------------
反射结构体 从它的地址创建了反射对象，然后修改结构体中的字段
反射结构体 将 typeOfT 设置为结构体的类型， 然后以直白的方法调用遍历其字段
------------------------------------------------------------------
可设置性的要点： T 的字段名必须大写（已导出）， 因为只有已导出的字段才是可设置的
------------------------------------------------------------------
反射法则如下：

	从接口值可反射出反射对象。
	从反射对象可反射出接口值。
	要修改反射对象，其值必须可设置
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"reflect"
)

type Entity struct {
	A int
	B string
}

func main() {

	var entity = Entity{ 22, "jack" }

	// 通过 反射地址，创建反射对象
	e := reflect.ValueOf(entity).Elem()

	// 获取类型
	typofT := e.Type()

	for i := 0; i < e.NumField(); i++ {
		field := e.Field(i)

		fmt.Printf("%d: %s %s = %v \n", i, typofT.Field(i).Name, field.Type(), field.Interface())
	}

	// output
	// --------------------------------------------------
	// 0: A int = 22
	// 1: B string = jack

	// 如果不通过地址反射，下面设置属性值将失败；因为 t 的字段不可设置
	// e = reflect.ValueOf(entity).Elem()

	// 通过反射，重新设置实体属性值
	e.Field(0).SetInt(77)
	e.Field(1).SetString("Sunset Strip")

	fmt.Println("entity is now：", entity)

}