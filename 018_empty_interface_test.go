package go_hello_world

import (
	"fmt"
	"testing"
)

// 空接口用来断言
func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}

	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknown type")
}

func DoSomethingSwitch(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknown type")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(1)
	DoSomething("10")
	DoSomething(float32(2.33))
}

func TestEmptyInterfaceSwitch(t *testing.T) {
	DoSomethingSwitch(1)
	DoSomethingSwitch("10")
	DoSomethingSwitch(float32(2.33))
}
