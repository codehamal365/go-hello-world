package go_hello_world

// 接口最佳实践
// 倾向于使用更小的接口定义，很多接口只包行一个接口
// 较大的接口，可以包含很多较小的接口，组合而成

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func StoreData(reader Reader) error {
	// do something
	return nil
}
