package go_hello_world

import (
	"fmt"
	"testing"
)

// 不支持子类重写的多态形式
// struct 类型 方法等扩展和复用
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

// 这里不支持子类重写
func (d *Dog) Speak() {
	fmt.Print("dog")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("wang!")
}
