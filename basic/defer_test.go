package basic

import (
	"fmt"
	"testing"
)

func test() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

func testHaveName() (i int) {
	i = 0
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

func TestDefer(t *testing.T) {
	// return 0. 返回值并没有被修改，这是由于 Go 的返回机制决定的，执行 return 语句后，Go 会创建一个临时变量保存返回值，因此，defer 语句修改了局部变量 i，并没有修改返回值
	fmt.Println("return", test())

	// return 1. 对于有名返回值的函数，执行 return 语句时，并不会再创建临时变量保存，因此，defer 语句修改了 i，即对返回值产生了影响
	fmt.Println("return", testHaveName())
}
