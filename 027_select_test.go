package go_hello_world

import (
	"testing"
	"time"
)

// 多路选择和超时控制
// 见026——csp_channel_test.go

func TestSelect(t *testing.T) {
	// service 执行时间设置为500
	select {
	case ret := <-BufferedAsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
