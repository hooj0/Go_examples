// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-23
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
reader exercise
==================================================================
reader 练习
------------------------------------------------------------------
实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。
------------------------------------------------------------------*/
package main

import (
	"golang.org/x/tour/reader"
	"io"
)

type MyReader struct{
	c int8
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(bytes []byte) (int, error) {

	if r.c != 'a' && r.c != 'A' {
		return 0, io.EOF
	}

	bytes[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{'A'})
}
