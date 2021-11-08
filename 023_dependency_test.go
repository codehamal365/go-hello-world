package go_hello_world

/**
go未解决的依赖问题
1.同一环境下，不同项目使用同一包的不同版本
2.无法管理对包的特定版本的依赖
*/

/**
go1.5后 vendor目录被添到除了GOPATH和GOROOT之外的依赖目录查找的解决方案
在1.6之前，需要手动设置环境变量
查找依赖包路径的解决方案如下：
1.当前包下的vendor目录
2.向上级目录查找，直到找到src下的vendor目录
3.在GOPATH下查找依赖包
4.在GOROOT目录下查找
*/

/**
常用的依赖管理工具
1.godep github.com/tools/dodep
2.glide github.com/Masterminds/glide
3.dep github.com/golang/dep
*/

/**
glide详细用法，可参考相关文档
*/
