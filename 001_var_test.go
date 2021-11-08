package go_hello_world

import (
	"testing"
)

// 变量赋值 自动类型推断
func TestFibList(t *testing.T) {

	// 方式一：初始化
	var a, b = 1, 1

	// 方式二:
	c, d := 1, 2
	t.Log(c, d)

	// 方式三
	var e, f int = 1, 2
	t.Log(e, f)

	// 方式四
	var (
		a1 = 1
		a2 = 2
	)
	t.Log(a1, a2)

	for i := 0; i < 10; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = a + tmp
	}
}

func TestExchange(t *testing.T) {
	var (
		a = 1
		b = 2
	)
	a, b = b, a
	t.Log(a, b)
}
