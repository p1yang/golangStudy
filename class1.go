// 变量和常量
package main

import "fmt"

// 声明全局变量声明可以不使用
var s1 string
var _ int
var _ bool

// 常量，不可修改
const pi = 3.1415

// 常量声明时如果某一行没有复制，那他与上一行值相同
const (
	n1 = 100
	n2
	n3 = 200
	n4
)

// iota 常量计数器，在const出现时被重置为0，const中每新增`一行`常数声明，iota增加1
const (
	a1 = iota
	a2
	a3
)

const (
	b1, b2 = iota + 1, iota + 2
	b3, b4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GM = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//非全局变量声明必须使用
	//abc,_,_123,a123
	var name string
	name = "zhangsan"
	fmt.Println(name)

	//类型推导
	var sex = "男"
	fmt.Println(sex)

	//短变量声明
	email := "admin@qq.com"
	fmt.Println(email)

	//匿名变量，不需要的变量，但需要接收，使用`_`
	_, a := "d", "s"
	fmt.Println(a)

	fmt.Println(pi)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
}
