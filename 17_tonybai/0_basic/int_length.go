package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 无符号整形的长度
	// 由于这三个类型的长度是平台相关的，所以我们在编写有移植性要求的代码时，千万不要强依赖这些类型的长度。
	// 如果你不知道这三个类型在目标运行平台上的长度，可以通过 unsafe 包提供的 SizeOf 函数来获取，比如在 x86-64 平台上，它们的长度均为 8：
	var a, b = int(5), uint(6)
	var p uintptr = 0x12345678
	fmt.Println("signed integer a's length is", unsafe.Sizeof(a))   // 8
	fmt.Println("unsigned integer b's length is", unsafe.Sizeof(b)) // 8
	fmt.Println("uintptr's length is", unsafe.Sizeof(p))            // 8

	// 整形溢出

	var s int8 = 127
	s += 1 // 预期128，实际结果-128
	fmt.Println("s is", s)

	var u uint8 = 1
	u -= 2 // 预期-1，实际结果255
	fmt.Println("u is", u)

	// 进制输出

	var d int8 = 59
	fmt.Printf("%b\n", d) //输出二进制：111011
	fmt.Printf("%d\n", d) //输出十进制：59
	fmt.Printf("%o\n", d) //输出八进制：73
	fmt.Printf("%O\n", d) //输出八进制(带0o前缀)：0o73
	fmt.Printf("%x\n", d) //输出十六进制(小写)：3b
	fmt.Printf("%X\n", d) //输出十六进制(大写)：3B
}
