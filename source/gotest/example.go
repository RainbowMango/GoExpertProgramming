package gotest

import "fmt"

// SayHello 打印一行字符串
func SayHello() {
	fmt.Println("Hello World")
}

// SayGoodbye 打印两行字符串
func SayGoodbye() {
	fmt.Println("Hello,")
	fmt.Println("goodbye")
}

// PrintNames 打印学生姓名
func PrintNames() {
	students := make(map[int]string, 4)
	students[1] = "Jim"
	students[2] = "Bob"
	students[3] = "Tom"
	students[4] = "Sue"
	for _, value := range students {
		fmt.Println(value)
	}
}
