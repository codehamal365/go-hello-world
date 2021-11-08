package go_hello_world

import "testing"

// map初始化
func TestMapInit(t *testing.T) {
	// 初始化方式一
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[3])
	t.Log("len m1: ", len(m1))

	// 初始化方式二
	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2[1])               // 0
	t.Log("len m2: ", len(m2)) // 1

	// 初始化方式三
	m3 := make(map[int]int, 10)
	t.Log("len m3: ", len(m3))
}

func TestAccessKeyNotExisting(t *testing.T) {
	m1 := map[int]int{}
	// 不存在时是0
	t.Log(m1[1])
	// 设置为0也是0
	m1[2] = 0
	t.Log(m1[2])
	// 如何判断key值是否存在
	if v, ok := m1[3]; ok {
		t.Logf("key is exist, value is %d", v)
	} else {
		t.Log("key is not exist")
	}
}

func TestRangeMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 4, 3: 9}
	for key, val := range m {
		t.Log(key, val)
	}
}
