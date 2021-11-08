package go_hello_world

import "testing"

// 通过接口实现多态
type JavaProgrammer struct {
}

func (javaProgrammer *JavaProgrammer) WriteHelloWorld() string {
	return "java programmer hello world"
}

func writeHelloWorld(p Programmer) string {
	return p.WriteHelloWorld()
}

func TestPolymorphism(t *testing.T) {
	goProgrammer := new(GoProgrammer)
	javaProgrammer := new(JavaProgrammer)

	t.Log(writeHelloWorld(goProgrammer))
	t.Logf(writeHelloWorld(javaProgrammer))
}
