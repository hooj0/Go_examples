// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-30
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
等价二叉树
==================================================================
可以用多种不同的二叉树的叶子节点存储相同的数列值。
例如，这里有两个二叉树保存了序列 1，1，2，3，5，8，13。

用于检查两个二叉树是否存储了相同的序列的函数在多数语言中都是相当复杂的。
这里将使用 Go 的并发和 channel 来编写一个简单的解法。

这个例子使用了 tree 包，定义了类型：

type Tree struct {
Left *Tree
Value int
Right *Tree }
1. 实现 Walk 函数。
2. 测试 Walk 函数。

函数 tree.New(k) 构造了一个随机结构的二叉树，保存了值 k，2k，3k，…，10k。
创建一个新的 channel ch 并且对其进行步进：

go Walk(tree.New(1), ch) 然后从 channel 中读取并且打印 10 个值。应当是值 1，2，3，…，10。
3. 用 Walk 实现 Same 函数来检测是否 t1 和 t2 存储了相同的值。
4. 测试 Same 函数。

Same(tree.New(1), tree.New(1)) 应当返回 true，而 Same(tree.New(1),
tree.New(2)) 应当返回 false。
------------------------------------------------------------------*/
package main

import "golang.org/x/tour/tree"


// Walk 递归 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}


	Walk(t.Left, ch)	// 分别将左边和右边的节点进行递归
	ch <- t.Value		// 将值发送到通道
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)	// 构建通道
	ch2 := make(chan int)

	go func() {				// 递归tree1
		Walk(t1, ch1)
		println("tree1: ", t1.Value)
		ch1 <- 0
	}()

	go func() {
		Walk(t2, ch2)		// 递归tree2
		println("tree1: ", t2.Value)
		ch2 <- 0
	}()

	for {
		t1 := <- ch1
		t2 := <- ch2

		println("t1: ", t1, ", t2: ", t2)
		if t1 == 0 && t2 == 0 {	// 递归结束，协程执行完毕
			return true
		}

		if t1 == t2 {			// 两个节点相等，继续递归
			continue
		} else {
			return false		// 两个节点不等，直接返回，结束程序
		}
	}

	return true
}

func main() {

	ch := make(chan int)

	go func() {

		Walk(tree.New(1), ch)
		ch <- 0 	// 退出标记
	}()

	for {
		t := <- ch
		if t == 0 {	// 接收上面的 0
			break
		}

		println(t)
	}

	println(Same(tree.New(1), tree.New(2)))
}
