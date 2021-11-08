package go_hello_world

import (
	"strconv"
	"strings"
	"testing"
)

/*
  string是数值类型，不是引用或指针类型
  string是只读的byte slice, len函数可以获取长度
  string的byte数值可以存放任何数据
*/

func TestString_(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(s)
	t.Log(len(s))

	// string不可变
	// s[1] = 's'

	r := []rune(s)
	t.Logf("中 unicode %x", r[0])
	t.Logf("中 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	// range 字符串自动转化为rune
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}

// 字符串函数测试
// strings
func TestStringFunc(t *testing.T) {
	// split
	s := "a:b:c"
	split := strings.Split(s, ":")
	t.Log(split)
	// join
	join := strings.Join(split, "-")
	t.Log(join)
}

// 字符串到int等转换
func TestConv(t *testing.T) {
	itoa := strconv.Itoa(10)
	t.Logf("%[1]T %[1]s", itoa)

	atoi, _ := strconv.Atoi("10")
	t.Logf("%[1]T %[1]d", atoi)
}
