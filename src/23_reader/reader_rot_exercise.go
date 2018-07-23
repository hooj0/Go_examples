// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-23
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
reader rot
==================================================================
有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。
------------------------------------------------------------------
gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）
并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。
------------------------------------------------------------------
编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，
通过应用 rot13 代换密码对数据流进行修改。
rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。
--------------------------------------------------------------------

rot13 代换密码
--------------------------------------------------------------------
套用ROT13到一段文字上仅仅只需要检查字元字母顺序并取代它在13位之后的对应字母 ，
有需要超过时则重新绕回26英文字母开头即可[2]。A换成N、B换成O、依此类推到M换成Z，
然后序列反转：N换成A、O换成B、最后Z换成M。只有这些出现在英文字母里头的字元受影响；
数字、符号、空白字元以及所有其他字元都不变。
因为只有在英文字母表里头只有26个，并且26=2×13，ROT13函数是它自己的逆反
https://blog.csdn.net/qq_27818541/article/details/54379030
------------------------------------------------------------------*/


package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}


func (rot rot13Reader) Read(bytes []byte) (int, error) {
	length, err := rot.r.Read(bytes)

	for i, b := range bytes[:length] {
		fmt.Println("i: ", i, ", b: ", b)
		bytes[i] = rot13(b)
		fmt.Println("i: ", i, ", b: ", bytes[i])
		fmt.Println("")
	}

	return length, err
}


// 转换byte  前进13位/后退13位
func rot13(b byte) byte {
    switch {
		case 'A' <= b && b <= 'M':
			b = b + 13
		case 'M' < b && b <= 'Z':
			b = b - 13
		case 'a' <= b && b <= 'm':
			b = b + 13
		case 'm' < b && b <= 'z':
			b = b - 13
    }
    return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
