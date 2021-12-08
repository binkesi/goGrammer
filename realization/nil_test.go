/*
接口(interface) 是对非接口值(例如指针，struct等)的封装，内部实现包含 2 个字段，类型 T 和 值 V。一个接口等于 nil，当且仅当 T 和 V 处于 unset 状态（T=nil，V is unset）
两个接口值比较时，会先比较 T，再比较 V。
接口值与非接口值比较时，会先将非接口值尝试转换为接口值，再比较。
*/
package realization

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	var p *int = nil
	var i interface{} = p
	fmt.Println(p == nil)
	fmt.Println(i == p)
	// 将一个 nil 非接口值 p 赋值给接口 i，此时，i 的内部字段为(T=*int, V=nil)
	// i 与 p 作比较时，将 p 转换为接口后再比较，因此 i == p
	// 当 i 与 nil 比较时，会将 nil 转换为接口 (T=nil, V=nil)，与i (T=*int, V=nil) 不相等
	fmt.Println(i == nil)
}
