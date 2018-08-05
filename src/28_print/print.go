// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-08-05
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
print 打印
==================================================================
fmt 包中的常用打印函数，且函数名首字母均为大写：如 fmt.Printf、fmt.Fprintf，fmt.Sprintf 等
------------------------------------------------------------------
Sprintf 字符串函数（Sprintf 等）会返回一个字符串，而非填充给定的缓冲区。
Fprint 	接受任何实现了 io.Writer 接口的对象作为第一个实参；如 os.Stdout 与 os.Stderr
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"os"
	"time"
)

type E struct {
	a int
	b float64
	c string
}

// 控制自定义类型的格式，为该类型定义一个具有 String() string 签名的方法。
// 重写类型 T 的String方法，定义自定义输出格式，可进行如下操作。
func (t *E) String() string {

	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

type MyString string

// 请勿通过调用 Sprintf 来构造 String 方法，因为它会无限递归你的的 String 方法
func (m MyString) String() string {

	return fmt.Sprintf("MyString=%s", m) 		// 错误：会无限递归
	return fmt.Sprintf("MyString=%s", string(m)) // 可以：注意转换
}

func main() {

	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	fmt.Printf("%v\n", time.Now())  // 或只用 fmt.Println(timeZone)

	// ----------------------------------------------------------
	// Hello 23
	// Hello 23
	// Hello 23
	// Hello 23
	// 2018-08-05 20:32:03.798148 +0800 CST m=+0.000000001


	type T struct {
		a int
		b float64
		c string
	}
	t := &T{ 7, -2.35, "abc\tdef" }
	fmt.Printf("%v\n", t)			// &{7 -2.35 abc	def}
	// 数输出带键值、属性的数据
	fmt.Printf("%+v\n", t)			// &{a:7 b:-2.35 c:abc	def}
	// 输出带键值属性和数据类型的字符串
	fmt.Printf("%#v\n", t)			// &main.T{a:7, b:-2.35, c:"abc\tdef"}
	fmt.Printf("%#v\n", time.Now())	// time.Time{wall:0xbed1d95bca0ee890, ext:1, loc:(*time.Location)(0x546e80)}


	// 输出数据类型
	fmt.Printf("%T\n", time.Now())	// time.Time

	// 输出带引号的字符串
	fmt.Printf("%q", "hello")		// "hello"

	// 重写 String方法后，输出如下
	e := &E{ 7, -2.35, "abc\tdef" }
	fmt.Printf("%v\n", e)			// "hello"7/-2.35/"abc\tdef"

}
