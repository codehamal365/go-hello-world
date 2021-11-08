package go_hello_world

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

// if
func TestMultiSec(t *testing.T) {
	if result, err := test(2); err == nil {
		t.Log("nil result", result)
	} else {
		t.Log("!nil result", result)
	}
}

func TestIfElse(t *testing.T) {
	var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	t.Log(prompt)
}

func TestIfElseMulti(t *testing.T) {
	var first = 10
	var cond int

	if first <= 0 {

		t.Logf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {

		t.Logf("first is between 0 and 5\n")
	} else {

		t.Logf("first is 5 or greater\n")
	}
	if cond = 5; cond > 10 {

		t.Logf("cond is greater than 10\n")
	} else {

		t.Logf("cond is not greater than 10\n")
	}
}

func test(a int) (int, error) {
	if a > 1 {
		return a, errors.New("e")
	}
	return a * a, nil
}

// switch 默认自带break
// 条件表达式不限制为常量或者整数
// 单个case中，可以出现多个结果选项，使用逗号分隔
// 不需要break来退出一个case
// 可以使用fallthrough继续执行下面的case
// 可以不设定switch后的条件表达式，整个switch结构相当于多个if else
// switch
func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X")
	case "linux":
		t.Log("linux")
	default:
		t.Logf("%s.", os)
	}
}

// switch if else 写法
func TestSwitch2(t *testing.T) {
	var num int
	switch {
	case 0 <= num && num <= 3:
		t.Log("0-3")
	case 4 <= num && num <= 6:
		t.Log("0-3")
	}
}

func TestSwitch3(t *testing.T) {

	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log(i, "0 or 2")
		case 1, 3:
			t.Log(i, "1 or 3")
		default:
			t.Log(i, "4")
		}
	}
}
