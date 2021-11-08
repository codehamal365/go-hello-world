package go_hello_world

import "testing"

// 初始化
func TestArrayInit(t *testing.T) {
	// 初始化数组 方式1
	var arr [3]int
	arr[0] = 1
	t.Log(arr[0], arr[1], arr[2])
	// 方式二
	arr1 := [4]int{1, 2, 3, 4}
	t.Log(arr1)
	// 方式三
	arr2 := [...]int{1, 2, 3, 5}
	t.Log(arr2)
}

// 数组遍历
func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 5, 7}
	// fori
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	// for range
	for index, element := range arr {
		t.Log(index, element)
	}
	// ignore index
	for _, element := range arr {
		t.Log(element)
	}
}

// 数组截取
// a[开始索引（包含）, 结束索引（不包含）] --- 含头不含尾
func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	arrSec1 := arr[:]
	t.Log(arrSec1)
}
