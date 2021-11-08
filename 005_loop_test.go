package go_hello_world

import "testing"

// while
func TestWhileLoop(t *testing.T) {
	var n = 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

// break
func TestLoopBreak(t *testing.T) {
	var n = 10
	for {
		if n < 5 {
			break
		}
		n--
	}
}

// continue
func TestLoopContinue(t *testing.T) {
	var n = 10
	for {
		n--
		if n == 5 {
			continue
		} else if n == 3 {
			break
		}

	}
}

func TestLoopLabel(t *testing.T) {

label:
	for a := 1; a <= 10; a++ {
		t.Log(a)
		for b := 1; b <= 10; b++ {
			t.Log(b)
			if a*b == 20 {
				break label
			}
		}
	}
}
