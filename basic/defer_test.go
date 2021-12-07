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

func TestDefer(t *testing.T) {
	fmt.Println("return", test())
}
