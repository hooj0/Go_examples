// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-12
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
test
==================================================================
测试框架 轻量级的测试框架，它的实现代码由 go test 命令和 testing 包构成
------------------------------------------------------------------
测试框架 通过创建一个名字以 _test.go 结尾的，
		包含名为 TestXXX 且签名为 func (t *testing.T) 函数的文件来编写测试。

测试框架 测试框架会运行每一个TestXXX(t *testing.T) 的函数
测试框架 TestXXX(t *testing.T) 函数调用了像 t.Error 或 t.Fail 这样表示失败的函数，此测试即表示失败。
------------------------------------------------------------------
测试框架 使用 go test 命令运行该测试：
	$ go test github.com/user/util
若在包目录下运行 go 工具，也可以忽略包路径
	$ go test
------------------------------------------------------------------*/
package util

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		input,  output string
	} {
		{"abc", "cba" },
		{"Hello, world", "dlrow ,olleH" },
		{"Hello, 世界", "界世 ,olleH" },
		{"", "" },
	}

	for _, c := range cases {
		result := Reverse(c.input)

		if result != c.output {
			t.Errorf("Reverse(%q) == %q, output %q", c.input, result, c.output)
		}
	}
}
