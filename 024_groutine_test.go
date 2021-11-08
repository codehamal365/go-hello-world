package go_hello_world

import (
	"fmt"
	"testing"
	"time"
)

/**
 thread vs groutine
1. 创建时默认的stack大小
-- jdk5后java thread stack 默认大小 1M
-- groutine的stack初始化大小为2k
2. 和KSE(kernel space entity)的对应关系
-- java thread是1:1
-- groutine是m:n
*/
func TestGroutine(t *testing.T) {

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}
