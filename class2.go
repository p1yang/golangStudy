// 整型
package main

import "fmt"

func main() {
	//uint和int根据操作系统的位数而改变。
	//常见有 int8 int16 int32 int64还有其对应的无符号uint
	var A1, A2 int8 = -128, 127
	fmt.Println(A1, A2)
	fmt.Println(A1-1, A2+1)

	var B1, B2 int16 = -32768, 32767
	fmt.Println(B1, B2)
	fmt.Println(B1-1, B1+1)

	var C1, C2 int32 = -2147483648, 2147483647
	fmt.Println(C1, C2)
	fmt.Println(C1-1, C1+1)

	var D1, D2 int64 = -9223372036854775808, 9223372036854775807
	fmt.Println(D1, D2)
	fmt.Println(D1-1, D1+1)

	//uint
	var a1, a2 uint8 = 0, 255
	fmt.Println(a1, a2)
	fmt.Println(a1-1, a2+1)

	var b1, b2 uint16 = 0, 65535
	fmt.Println(b1, b2)
	fmt.Println(b1-1, b2+1)

	var c1, c2 uint32 = 0, 4294967295
	fmt.Println(c1, c2)
	fmt.Println(c1-1, c2+1)

	var d1, d2 uint64 = 0, 18446744073709551615
	fmt.Println(d1, d2)
	fmt.Println(d1-1, d2+1)

}
