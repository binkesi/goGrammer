/*
比较两个interface
Go 语言中，interface 的内部实现包含了 2 个字段，类型 T 和 值 V，
interface 可以使用 == 或 != 比较。2 个 interface 相等有以下 2 种情况：
1.两个 interface 均等于 nil（此时 V 和 T 都处于 unset 状态）
2.类型 T 相同，且对应的值 V 相等。
*/
package realization

import (
	"fmt"
	"testing"
)

type Stu struct {
	Name string
}

type StuInt interface{}

var stu1, stu2 StuInt = &Stu{Name: "Tom"}, &Stu{Name: "Tom"}
var stu3, stu4 StuInt = Stu{Name: "Tom"}, Stu{Name: "Tom"}

func TestInterface(t *testing.T) {
	fmt.Println(stu1 == stu2)
	fmt.Println(stu3 == stu4)
}
