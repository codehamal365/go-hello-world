package go_hello_world

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
1.可以返回多个值
2.所有参数都是值传递：slice，map，channel会有传引用的错觉
3.函数可以作为变量的值
4.函数可以作为参数和返回值
*/
// 多值返回
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFN(t *testing.T) {
	_, va := returnMultiValues()
	t.Log(va)
}

// 对相同类型入参和出参的函数进行包装
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Print("time spent:", time.Since(start))
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 10)
	return op
}

func TestTimeSpentWrap(t *testing.T) {
	res := timeSpent(slowFunc)(10)
	t.Log(res)
}

// 可变参数
func sum(ops ...int) int {
	var res int
	for _, value := range ops {
		res += value
	}
	return res
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3))
}

// defer 延迟执行函数 不管程序是否panic都会执行
// defer 执行顺序 栈顺序 先进后出
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear resources === 1")
	}()
	defer func() {
		t.Log("clear resources === 2")
		//recover()
	}()
	t.Log("started")
	panic("fatal error")
	//t.Log("end")
}
