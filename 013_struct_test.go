package go_hello_world

import (
	"fmt"
	"testing"
	"unsafe"
)

// 对象数据、行为的封装
type Employee struct {
	Id   string
	Name string
	Age  int
}

// 初始化struct
func TestStructInit(t *testing.T) {
	// 方式一
	employee1 := Employee{"1", "zhangsan", 18}
	// 方式二
	employee2 := Employee{Id: "2", Name: "zhangsan", Age: 18}
	// 方式三
	employee3 := new(Employee)
	employee3.Id = "3"
	t.Log(employee1, employee2, employee3)
	t.Logf("employee1 %T", employee1)
	t.Logf("employee3 %T", employee3)
}

// struct 添加行为
// 第一种定义方式在实例对象方法被调用时，实例成员会进行赋值
// 也就是说对应实例对象调用时会进行值拷贝
func (e Employee) String() string {
	fmt.Printf("Access is %x", unsafe.Pointer(&e.Name))
	fmt.Println()
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 通常情况下为了避免内存拷贝我们使用第二种方式
func (e *Employee) StringPtr() string {
	fmt.Printf("Access is %x", unsafe.Pointer(&e.Name))
	fmt.Println()
	return fmt.Sprintf("Id:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

// go里不管通过实例还是指针 表述一样
func TestStructOperation(t *testing.T) {
	e := Employee{"0", "tom", 28}
	fmt.Printf("Access is %x", unsafe.Pointer(&e.Name))
	t.Log(e.String())
	t.Log(e.StringPtr())

	//e1:= &Employee{"0","tom",28}
	//t.Log(e1.String())
	//t.Log(e1.StringPtr())
}
