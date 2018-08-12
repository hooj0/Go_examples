package main

// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-12
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
main 函数主程序
==================================================================
main 函数主程序 是唯一的入口函数，是程序的入口
------------------------------------------------------------------
main 函数主程序 创建一个工程目录
	$ mkdir $GOPATH/src/github.com/user/hello

main 函数主程序 用 go 工具构建并安装此程序：
	$ go install github.com/user/hello

也可以直接在文件所在目录中运行命令：go install
		$ cd $GOPATH/src/github.com/user/hello
		$ go install
------------------------------------------------------------------
此命令会构建 hello 命令，产生一个可执行的二进制文件。
接着它会将该二进制文件作为 hello（在 Windows 下则为 hello.exe）
安装到工作空间的 bin 目录中。

go 工具只有在发生错误时才会打印输出，因此若这些命令没有产生输出， 就表明执行成功了。
------------------------------------------------------------------
main 函数主程序 当执行 install 后，可以执行命令运行程序
	$ $GOPATH/bin/hello
或者
	$ cd $GOPATH/bin/
	$ hello

若你已经将 $GOPATH/bin 添加到 PATH 中了，只需输入该二进制文件名即可：
	$ hello
------------------------------------------------------------------*/

import (
	"fmt"
	"github.com/hooj0/util"
)



func main() {

	fmt.Println("hello ", util.Reverse("world!"))
}
