package gotest_test

import (
	"testing"

	"github.com/rainbowmango/goexpertprogramming/source/gotest"
)

func benchSub1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithoutAlloc()
	}
}

func benchSub2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithoutAlloc()
	}
}

func benchSub3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithoutAlloc()
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("A=1", benchSub1)
	b.Run("A=2", benchSub2)
	b.Run("B=1", benchSub3)
}
