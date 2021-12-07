package basic

import (
	"fmt"
	"testing"
)

type StuType int32

const (
	Type1 StuType = iota
	Type2
	Type3
	Type4
)

func TestConst(t *testing.T) {
	fmt.Println(Type1, Type2, Type3, Type4) // 0, 1, 2, 3
}
