package go_hello_world

import (
	"errors"
	"testing"
)

// go异常机制 相比其他语言 go没有异常处理机制
// error类型实现了 error接口
// 可以通过errors.New来快速创建错误实例
// 一般处理异常情况是快速return 如果err不为空直接return

var NotLessThanTwoError = errors.New("n should not be less than 2")

var LessThanOneHundredError = errors.New("n should be less than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, NotLessThanTwoError
	}
	if n > 100 {
		return nil, LessThanOneHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if fibonacci, err := GetFibonacci(-10); err != nil {
		if err == NotLessThanTwoError {
			t.Log("it is less")
		}
		t.Error(err)
	} else {
		t.Log(fibonacci)
	}
}
