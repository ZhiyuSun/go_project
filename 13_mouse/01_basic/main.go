package main

import (
	"bytes"
	"fmt"
	"reflect"
)

type data struct {

}

type Person struct {
	Name string
	Sexual string
	Age int
}

func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}



type Country struct {
	Name string
}

type City struct {
	Name string
}

type Stringable interface {
	ToString() string
}
func (c Country) ToString() string {
	return "Country = " + c.Name
}
func (c City) ToString() string{
	return "City = " + c.Name
}

func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}



type Shape interface {
	Sides() int
	Area() int
}
type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}


func main()  {
	// 01. Go编程模式：切片、接口、时间和性能

	// Slice
	fmt.Println("=====case1=====")
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100
	fmt.Println(foo)
	bar  := foo[1:4]
	bar[1] = 99
	fmt.Println(bar)
	fmt.Println(foo) // [0 0 99 42 100]

	fmt.Println("\n=====case2=====")
	a := make([]int, 32)
	fmt.Println(len(a)) // 32
	fmt.Println(cap(a)) // 32
	b := a[1:16]
	fmt.Println(len(b)) // 15
	fmt.Println(cap(b)) // 31
	a = append(a, 1)
	fmt.Println(len(a)) // 33
	fmt.Println(cap(a)) // 64
	a[2] = 42
	fmt.Println(len(b)) // 15
	fmt.Println(cap(b)) // 31
	fmt.Println(a) // [0 0 42 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
	fmt.Println(b) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	// append()这个函数在 cap 不够用的时候，就会重新分配内存以扩大容量，如果够用，就不会重新分配内存了！

	fmt.Println("\n=====case3=====")
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')

	dir1 := path[:sepIndex]
	//dir1 := path[:sepIndex:sepIndex]  // 指定了dir1的容量，append后会重新分配内存
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB

	fmt.Println(len(dir1)) // 4
	fmt.Println(cap(dir1)) // 14
	fmt.Println(len(dir2)) // 9
	fmt.Println(cap(dir2)) // 9

	dir1 = append(dir1,"suffix"...)

	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB

	fmt.Println(len(dir1)) // 10
	fmt.Println(cap(dir1)) // 14
	fmt.Println(len(dir2)) // 9
	fmt.Println(cap(dir2)) // 9

	fmt.Println("\n=====case4 深度比较=====")
	// 当我们复制一个对象时，这个对象可以是内建数据类型、数组、结构体、Map……在复制结构体的时候，
	//如果我们需要比较两个结构体中的数据是否相同，就要使用深度比较，而不只是简单地做浅度比较。这里需要使用到反射 reflect.DeepEqual()
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:",reflect.DeepEqual(v1,v2))
	//prints: v1 == v2: true

	m1 := map[string]string{"one": "a","two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:",reflect.DeepEqual(m1, m2))
	//prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:",reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true

	fmt.Println("\n=====case5 接口编程=====")
	var p = Person{
		Name: "Hao Chen",
		Sexual: "Male",
		Age: 44,
	}
	PrintPerson(&p) // 函数
	p.Print() // 成员函数
	//  Go 语言中，使用“成员函数”的方式叫“Receiver”，这种方式是一种封装，因为 PrintPerson()本来就是和 Person强耦合的，所以理应放在一起。
	// 更重要的是，这种方式可以进行接口编程，对于接口编程来说，也就是一种抽象，主要是用在“多态”


	fmt.Println("\n=====case6 接口编程=====")
	d1 := Country {"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)

	fmt.Println("\n=====case6 接口完整性检查=====")
	//var _ Shape = (*Square)(nil)
	s := Square{len: 5}
	fmt.Printf("%d\n",s.Sides())
}
