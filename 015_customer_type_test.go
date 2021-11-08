package go_hello_world

import (
	"fmt"
	"testing"
	"time"
)

// 自定义类型 参加 012_function_test
type FuncWrap func(op int) int

func customerTimeSpent(inner FuncWrap) FuncWrap {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Print("time spent:", time.Since(start))
		return ret
	}
}

func TestFuncWrap(t *testing.T) {
	t.Log(customerTimeSpent(func(op int) int {
		return op
	})(2))
}
