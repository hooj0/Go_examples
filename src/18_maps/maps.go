// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
map
==================================================================
映射 将键映射到值，键值对数据结构类型
------------------------------------------------------------------

------------------------------------------------------------------
映射 的初始默认值（零值） 为nil，nil既没有键，也不能添加键
映射 可以利用make创建，make初始化作为备用
映射 键值可以是任何有效的数据类型
映射 键不能重复，重复键将覆盖之前的数据
映射 值不可以设置为 nil
------------------------------------------------------------------*/

package main

import "fmt"

// 声明一个结构体数据模型
type Vertex struct {
	email string
	age int
}

// 声明一个map对象，键为string类型，值为Vertex类型
var maps map[string]Vertex

func main() {

	// 初级初始化maps
	maps = make(map[string]Vertex)

	// 设置值
	maps["jack"] = Vertex{"jack@qq.com", 22}
	maps["jason"] = Vertex{"jason@qq.com", 34}

	// 取值
	fmt.Printf("jason: %v \n", maps["jason"])	// jason: {jason@qq.com 34}
	fmt.Printf("tom: %v \n", maps["tom"])		// tom: { 0}

	fmt.Println("maps: ", maps) // maps:  map[jack:{jack@qq.com 22} jason:{jason@qq.com 34}]
}
