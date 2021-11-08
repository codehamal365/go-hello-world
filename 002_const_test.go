package go_hello_world

import "testing"

// 常量赋值
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Open = 1 << iota
	_
	Close
	Pending
)

const (
	Readable = 1 + iota
	Writable
	Executable
)

func TestIota(t *testing.T) {
	// 1 2
	t.Log(Monday, Tuesday)

	// 1 4
	t.Log(Open, Close)

	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
