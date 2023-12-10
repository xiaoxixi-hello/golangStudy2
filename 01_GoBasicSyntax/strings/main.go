package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello")
	fmt.Println(`这是一行
 这是第二行 完全打印 包括空格等`)

	//	字符串长度
	println(len("你好"))                      // 输出6
	println(utf8.RuneCountInString("你好"))   // 输出2
	println(utf8.RuneCountInString("你好ab")) // 输出4

	// 遇到计算字符个数，比如用户名多长等
	//	使用 utf8.RuneCountInString

	//	 字符串与数字不能通过+拼接，当遇到需要拼接是如下
	s := fmt.Sprintf("%s和%f", "你好", 0.1234)
	fmt.Println(s)
}
