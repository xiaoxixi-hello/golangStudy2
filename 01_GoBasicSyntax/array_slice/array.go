package main

import "fmt"

// 数组
func main() {
	//	 直接初始化一个三个元素的数组，里面的元素不足就以0值补充
	a1 := [3]int{2, 3, 5}
	fmt.Printf("a1:%v, len:%d, cap: %d\n", a1, len(a1), cap(a1))

	numArray := [3]int{1, 2}
	fmt.Println(numArray)

	//	 直接初始化一个三个元素的数组，里面的元素只都是0
	var a2 [3]int
	fmt.Printf("a1:%v, len:%d, cap: %d\n", a2, len(a2), cap(a2))
}
