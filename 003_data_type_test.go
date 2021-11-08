package go_hello_world

import "testing"

// bool

// string

// int int8 int16 int32 int64

// uint uint8 uint16 uint32 uint64 uintptr

// byte == uint8

// rune == int32

// float32 float64

// complex64 complex128

// 注意：不支持隐式数据类型转换
func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	// 转换报错
	// b = a
	b = int64(a)
	t.Log(b)
}

// 类型的预定义值
// math.MaxInt64 max.maxFloat64 math.MaxUint32

// 不支持指针运算 string是值类型，其默认初始化值是空字符串，而不是nil
func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	// aPtr = aPtr + 1 报错
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	// 空串
	t.Log(len(s))
	t.Log("=", s, "=")
}
