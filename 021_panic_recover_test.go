package go_hello_world

import (
	"errors"
	"fmt"
	"testing"
)

// panic 用于不可以恢复的错误
// panic 退出前会执行 defer 指定的内容

// os.Exit退出时不会调用defer指定函数
// os.Exit退出时不输出当前调用栈信息
func TestPanicVsExit(t *testing.T) {
	fmt.Println("start")
	//os.Exit(-1)

	defer func() {
		fmt.Println("finally")
	}()
	panic(errors.New("something wrong"))
}

// 通过 recover 恢复
// 注意：recover后虽然程序会正常执行，导致进程没有结束，但是会
// 形成僵尸服务进程，导致health check失效(服务并不可用)
// let it crash, daemon进程重启服务
func TestRecover(t *testing.T) {
	fmt.Println("start")
	//os.Exit(-1)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	panic(errors.New("something wrong"))
}
