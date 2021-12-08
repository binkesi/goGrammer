/*
Go 语言的局部变量分配在栈上还是堆上?
由编译器决定。Go 语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析(escape analysis)，
当发现变量的作用域没有超出函数范围，就可以在栈上，反之则必须分配在堆上。
*/

package realization

import (
	"fmt"
	"testing"
)

func foo() *int {
	v := 10
	return &v
}

func TestVar(t *testing.T) {
	m := foo()
	fmt.Println(*m)
}
