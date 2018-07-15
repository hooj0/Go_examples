// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
map mutating
==================================================================
map增删改查 可以对map进行添加元素、删除元素、修改元素、查询元素操作
------------------------------------------------------------------
语法：
	var maps = make(map[Key类型]Value类型)
	// 添加
	maps[key] = value
	// 删除
	delete(maps, key)
	// 修改
	maps[key] = value

	// 查询
	// 若 key 在 maps 中，ok 为 true ；否则， ok 为 false。
	// 若 key 不在映射中，那么 elem 是该映射元素类型的零值。
	v, ok := maps[key]
------------------------------------------------------------------
map可以作为一种数据词典的数据结构使用

------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	var maps = make(map[string]int)
	// 添加
	maps["a"] = 1

	// 删除
	delete(maps, "a")

	// 查询
	v, ok := maps["a"]
	// 由于已经删除，值=0，ok 为false，表示不存在
	fmt.Println("v: ", v, ", ok: ", ok)	// v:  0 , ok:  false

	// 修改
	maps["a"] = 2

	// 再次查询
	v, ok = maps["a"]

	fmt.Println("v: ", v, ", ok: ", ok)	// v:  2 , ok:  true

}
