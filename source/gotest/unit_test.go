package gotest_test

import (
	"flag"
	"testing"

	"github.com/rainbowmango/goexpertprogramming/source/gotest"
)

func TestAdd(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}

// TestArgs 用于演示如何解析-args参数
func TestArgs(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}

	argList := flag.Args() // flag.Args() 返回 -args 后面的所有参数，以切片表示，每个元素代表一个参数
	for _, arg := range argList {
		if arg == "cloud" {
			t.Log("Running in cloud.")
		} else {
			t.Log("Running in other mode.")
		}
	}
}
