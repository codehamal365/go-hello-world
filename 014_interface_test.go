package go_hello_world

import "testing"

// go 接口是非入侵性的，实现不依赖接口定义
// 所有接口的定义 可以 包含在接口使用包内
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (programmer *GoProgrammer) WriteHelloWorld() string {
	return "go programmer hello world"
}

func TestProgrammerClient(t *testing.T) {
	goProgrammer := new(GoProgrammer)
	t.Log(goProgrammer.WriteHelloWorld())
}

// 接口变量
func TestInterfaceVariable(t *testing.T) {
	var pro Programmer = &GoProgrammer{}
	t.Log(pro)
}
