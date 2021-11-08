package go_hello_world

import "testing"

// 数组和切片
// 容量是否可伸缩、是否可以比较

// slice初始化
func TestSliceInit(t *testing.T) {
	// 初始化方式一
	var slice []int
	t.Log(len(slice), cap(slice))
	slice = append(slice, 1)
	t.Log(len(slice), cap(slice))

	// 初始化方式二 初始化值
	slice1 := []int{1, 2, 3}
	t.Log(len(slice1), cap(slice1))

	// 初始化三 make
	slice2 := make([]int, 3, 5)
	t.Log(len(slice2), cap(slice2))
	t.Log(slice2[0], slice2[1], slice2[2])
	// 下面这行报错
	//t.Log(slice2[0],slice2[1],slice2[2],slice2[3])
	slice2 = append(slice2, 1)
	t.Log(slice2[0], slice2[1], slice2[2], slice2[3])
}

// 切片 len 和 cap 变化，当切片存储不够的时候，cap是2倍增长
func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, 1)
		t.Log(len(s), cap(s))
	}
}

// 切片共享存储结构
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2), &Q2)
	Q2_1 := year[3:6]
	t.Log(Q2, len(Q2_1), cap(Q2_1), &Q2_1)
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer), &summer)
	summer[0] = "Unknown"
	t.Log(Q2)
}
