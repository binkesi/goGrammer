/*
一个T类型的值可以调用为*T类型声明的方法，但是仅当此T的值是可寻址(addressable) 的情况下。
编译器在调用指针属主方法前，会自动取此T值的地址。因为不是任何T值都是可寻址的，所以并非任何T值都能够调用为类型*T声明的方法。
反过来，一个*T类型的值可以调用为类型T声明的方法，这是因为解引用指针总是合法的。
事实上，你可以认为对于每一个为类型 T 声明的方法，编译器都会为类型*T自动隐式声明一个同名和同签名的方法。

哪些值是不可寻址的呢？

字符串中的字节；
map 对象中的元素（slice 对象中的元素是可寻址的，slice的底层是数组）；
常量；
包级别的函数等。
*/

package realization

import (
	"fmt"
	"testing"
)

type T string

func (t *T) Hello() {
	fmt.Println("Hello!")
}

func TestPointer(t *testing.T) {
	var t1 T = "abc"
	t1.Hello()
	const t2 T = "ABC"
	// t2.Hello()  can't call Hello() for const t2.
}
