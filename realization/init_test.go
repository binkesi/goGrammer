// init() 函数没有入参和返回值，不能被其他函数调用，同一个包内多个 init() 函数的执行顺序不作保证。
// 一句话总结： import –> const –> var –> init() –> main()
// init()函数是包作用域

package realization

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init1:", a)
}

func init() {
	fmt.Println("init2:", b)
}

var a = 1

const b = 2

func TestInit(t *testing.T) {
	fmt.Println("Test init function.")
}
