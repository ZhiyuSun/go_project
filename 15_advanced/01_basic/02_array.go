package main

import (
	"fmt"
)

func main()  {
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6
	fmt.Println(a, b, c, d)

	var a1 = [...]int{1, 2, 3} // a 是一个数组
	var b1 = &a1                // b 是指向数组的指针

	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", b1)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(a1[0], a1[1])   // 打印数组的前2个元素
	fmt.Println(b1[0], b1[1])   // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range b1 {     // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}

	// 字符串数组
	//var s1 = [2]string{"hello", "world"}
	//var s2 = [...]string{"你好", "世界"}
	//var s3 = [...]string{1: "世界", 0: "你好", }


	// 结构体数组
	//var line1 [2]image.Point
	//var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	//var line3 = [...]image.Point{{0, 0}, {1, 1}}

	// 图像解码器数组(函数数组)
	//var decoder1 [2]func(io.Reader) (image.Image, error)
	//var decoder2 = [...]func(io.Reader) (image.Image, error){
	//	png.Decode,
	//	jpeg.Decode,
	//}

	// 接口数组
	//var unknown1 [2]interface{}
	//var unknown2 = [...]interface{}{123, "你好"}

	// 管道数组
	//var chanList = [2]chan int{}

	// 空数组
	var d [0]int       // 定义一个长度为0的数组
	var e = [0]int{}   // 定义一个长度为0的数组
	var f = [...]int{} // 定义一个长度为0的数组


}