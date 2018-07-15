// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
map exercise
==================================================================
map练习 统计一个字符串出现次数
------------------------------------------------------------------
利用make构建一个简单的 map[string]int 类型
string key 存储单词，int value 存储单词出现次数
利用 strings.Fields() 方法进行分词
迭代遍历 分词后的数组

在map中查找迭代的当前单词，若出现在map中就将次数加1，否则为第一次 应该为1
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	// 利用 strings.Fields() 方法进行分词
	words := strings.Fields(s)

	// 利用make构建一个简单的 map[string]int 类型
	wordCountMap := make(map[string]int)
	// 迭代遍历 分词后的数组
	for _, word := range words {

		// 在map中查找迭代的当前单词
		count, flag := wordCountMap[word]

		// 若出现在map中就将次数加1，否则为第一次 应该为1
		if flag {
			wordCountMap[word] = count + 1
		} else {
			wordCountMap[word] = 1
		}
	}

	return wordCountMap
}

func main() {
	wc.Test(WordCount)

	// strings.Fields 将字符串用空格分割成数组
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}
