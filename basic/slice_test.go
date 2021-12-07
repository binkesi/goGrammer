package basic

import (
	"fmt"
	"testing"
)

/*
array和slice两者最大的差异是数组长度一旦确定便不可修改，无法变长，无法改短。而切片则通过伸缩容量。且数组是值类型，把一个数组赋予给另一个数组时是发生值拷贝，而切片是指针类型，拷贝的是指针。
var a1 []int ---> 无长度定义，是Slice
var a2 [2]int ---> 有长度定义，是Array
var a3 [3]*int---> 有长度定义，是Array
*/

func SliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSlice(t *testing.T) {
	var a = []string{"stu", "tom", "t001"}
	var b = []string{"stu", "tom", "t001"}
	result := SliceEqual(a, b)
	fmt.Println(result)
}
