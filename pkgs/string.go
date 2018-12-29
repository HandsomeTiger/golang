package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"

	//strings.Contains() 判断是否包含 返回true or false
	fmt.Println(strings.Contains(s, "hello"), strings.Contains(s, "?"))
	//找到首次出现的索引 类似于 str_pos
	fmt.Println(strings.Index(s, "o"))
	//切割字符串 explode
	ss := "1#2#345"
	splitedStr := strings.Split(ss, "#")
	fmt.Println(splitedStr)
	//合并字符串 implode
	joinStr := strings.Join(splitedStr, "##")
	fmt.Println(joinStr)
	//判断是否有某个前缀用 hasPrefix, 是否有某个后缀 hasSuffix

}
