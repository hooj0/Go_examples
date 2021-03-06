// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-02
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
package
==================================================================
包 主要用于组织文件
------------------------------------------------------------------
每个 Go 程序都是由包构成的
程序从 main 包开始运行
在 Go 程序中，包没有被使用需要删除，否则无法编译运行
------------------------------------------------------------------
包 用于管理多个*.go文件，将多个文件放在一个文件夹中，可以视为一个package
包 里面还可以有包
包 里面的文件名不能相同，但包里面的子包中的文件名可以和父包中的文件相同，不同的包中的文件名称可以相同
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("rand number: ", rand.Intn(10))

	//fmt.Println("cell number: ", math.Ceil(rand.Float64()))
}
