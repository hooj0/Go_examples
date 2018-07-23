// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-23
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
reader
==================================================================
读取数据流 是io包中的接口，可以从数据流末尾进行读取数据
------------------------------------------------------------------
Reader：
	func (T) Read(b []byte) (n int, err error)
------------------------------------------------------------------
Read 方法读取byte字节数组，并返回填充的字节长度和错误信息（io.EOF）
Read 方法读取缓冲字节数组，如果没有错误，将返回 <nil>
Read 方法填充后的字节数组的长度为0，不代表读完
Read 方法填充后的字节长度为读到的真实字节长度
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	// 创建一个实现过Reader接口的对象
	var s = strings.NewReader("IO Reader examples")
	// 创建一个读取字节的数组缓冲区
	bytes := make([]byte, 8)
	// 死循环读取数据，直到 错误 终止
	for {

		length, err := s.Read(bytes)

		fmt.Printf("read bytes: %v, length: %v, error: %v \n", bytes, length, err)
		fmt.Printf("byte[:n]: %v \n", bytes[:length])

		if err == io.EOF {
			break
		}
	}
}

// read bytes: [73 79 32 82 101 97 100 101], length: 8, error: <nil>
// byte[:n]: [73 79 32 82 101 97 100 101]
// read bytes: [114 32 101 120 97 109 112 108], length: 8, error: <nil>
// byte[:n]: [114 32 101 120 97 109 112 108]
// read bytes: [101 115 101 120 97 109 112 108], length: 2, error: <nil>
// byte[:n]: [101 115]
// read bytes: [101 115 101 120 97 109 112 108], length: 0, error: EOF
// byte[:n]: []
