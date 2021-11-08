package go_hello_world

import "testing"

// 基本运算 + - * / % ++ -- 无前置++ --

// 比较运算 == != > < >= <=

// 逻辑运算 && || ！

// 位运算 &(同1为1) |(有1为1) ^(不同为1) >>(右移) <<(左移)
// 与其他编程语言的差异
// 右边为1左边清0，右边是0左边保留原值
// 1 &^ 0 -- 1
// 1 &^ 1 -- 0
// 0 &^ 1 -- 0
// 0 &^ 0 -- 0
func TestBitCompute(t *testing.T) {
	var a = 7           // 0111 0001
	a = a &^ Readable   // 去掉可读
	a = a &^ Writable   // 去掉可写
	a = a &^ Executable // 去掉可执行
	t.Log(a)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

// ==比较数组
// 相同维数且含有相同元素个数数组才能比较  每个元素都相等才相等
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{3, 4, 5}
	c := [...]int{1, 2, 4, 5}
	d := [...]int{1, 2, 4, 5}
	t.Log(a == b)
	// 报错
	// t.Log(a==c)
	t.Log(c == d)
}
