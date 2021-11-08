package go_hello_world

import (
	"fmt"
	"testing"
)

// map的值可以是方法，与go 的duck type接口方式一起，可以方便的实现单一方法对象的工厂模式

func TestMapWithFunctionValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

// go的内置集合中没有Set实现，可以用map[type]bool来实现
/**
  1. 元素的唯一性
  2. 添加元素、判断元素是否存在、删除元素、元素个数
*/
func TestMapForSet(t *testing.T) {
	set := make(map[string]bool) // New empty set
	set["Foo"] = true            // Add
	for k := range set {         // Loop
		t.Log(k)
	}
	delete(set, "Foo")   // Delete
	size := len(set)     // Size
	exists := set["Foo"] // Membership
	t.Log(size, exists)
}

// map的value值是布尔型，这会导致set多占用内存空间，解决这个问题，
// 则可以将其替换为空结构。在Go中，空结构通常不使用任何内存
// unsafe.Sizeof(struct{}{}) // 结果为 0
type void struct{}

func TestMapForSetOptimize(t *testing.T) {
	var member void

	set := make(map[string]void) // New empty set
	set["Foo"] = member          // Add
	for k := range set {         // Loop
		fmt.Println(k)
	}
	delete(set, "Foo")      // Delete
	size := len(set)        // Size
	_, exists := set["Foo"] // Membership
	t.Log(size, exists)
}

// golang-set 第三方库
