package go_hello_world

import "testing"
import cm "github.com/easierway/concurrent_map"

/**
1.基本复用单元模块、以首字母大写来表明可被包外访问
2.代码的package可用和所在目录不一致
3.同一目录下的go代码package要保持一致
*/

// GOPATH/src/packageName/....
// 引用 import "packageName/..."

/*
init方法
1.在main被执行前，所有依赖的package的init方法都会被执行
2.不同包的init方法按照导入的依赖关系决定执行顺序
3.每个包可用有多个init函数
4.包的每个源文件也可以有多个init函数
*/

/**
如何使用远程的package
ConcurrentMap for Go
强制从网络更新远程依赖
go get -u github.com/easierway/concurrent_map
自己本地提交代码，不要把src目录提交，否则go get 获取不到
*/

func TestConCurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
