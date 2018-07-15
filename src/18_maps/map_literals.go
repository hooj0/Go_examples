// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
map literals
==================================================================
映射文法 键名为必须项，不能为空或不填
------------------------------------------------------------------
映射文法 和struct文法类似，可以定向为指定属性赋值
映射文法 键必须存在且不能重复，值不能为nil
------------------------------------------------------------------*/
package main

import "fmt"

// 声明一个 复杂struct类型
type Vertex struct {
	Email, NickName string
}


// 声明一个简单类型的map
var simpleMap = map[string]int{
	"z": 11122,
	"a": 'a',
	"b": 'b',
}

// 声明一个复杂类型 map
var complexMap = map[string]Vertex {
	"jack": Vertex {
		"jack@qq.com",
		"jack chen",
	},
	// Vertex 类型可以省略
	"jason": {
		"jason@qq.com",
		"jack chen",
	},
	// 简单赋值，为指定属性设置值
	"tom": Vertex{ NickName: "Tom Cruise" },
	"rachel": Vertex{},
}

func main() {

	fmt.Println("simple map: ", simpleMap) // simple map:  map[z:11122 a:97 b:98]

	fmt.Println("simple map: ", complexMap)
	// simple map:  map[jack:{jack@qq.com jack chen} jason:{jason@qq.com jack chen} tom:{ Tom Cruise}]
}
