package main

import "fmt"

var name string
var (
	na   string
	age  int
	isok bool
)

func main() {
	var test ="你峨眉"

	na = "111"
	isok = false
	name = "你好"
	//test="好的"
	fmt.Println(na)
	fmt.Println(name)
	fmt.Println(test)
	//剪短变量申明 只能函数里申明
	 test1 := "好的"
	fmt.Println(test1)
}
