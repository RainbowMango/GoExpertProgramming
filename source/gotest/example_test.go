package gotest_test

import "github.com/rainbowmango/goexpertprogramming/source/gotest"

// 检测单行输出
func ExampleSayHello() {
	gotest.SayHello()
	// OutPut: Hello World
}

// 检测多行输出
func ExampleSayGoodbye() {
	gotest.SayGoodbye()
	// OutPut:
	// Hello,
	// goodbye
}

// 检测乱序输出
func ExamplePrintNames() {
	gotest.PrintNames()
	// Unordered output:
	// Jim
	// Bob
	// Tom
	// Sue
}
