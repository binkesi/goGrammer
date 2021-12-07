package basic

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
}

// 但如果结构体定义了 String() 方法，%v 和 %+v 都会调用 String() 覆盖默认值。
func TestStruct(t *testing.T) {
	fmt.Printf("%v\n", Student{"Tom"})  // {Tom}
	fmt.Printf("%+v\n", Student{"Tom"}) // {Name:Tom}
}
