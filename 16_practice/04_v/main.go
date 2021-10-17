package main

import "fmt"

type student struct {
	id   int32
	name string
}

func main() {
	a := &student{id: 1, name: "xiaoming"}

	fmt.Printf("a=%v	\n", a)
	fmt.Printf("a=%+v	\n", a)
	fmt.Printf("a=%#v	\n", a)

	var b []*student
	b = append(b, &student{id: 1, name: "xiaoming"})
	b = append(b, &student{id: 2, name: "xiaohong"})

	fmt.Printf("b=%v	\n", b)
	fmt.Printf("b=%+v	\n", b)
	fmt.Printf("b=%#v	\n", b)

	c := []int64{1,2,3}

	fmt.Printf("c=%v	\n", c)
	fmt.Printf("c=%+v	\n", c)
	fmt.Printf("c=%#v	\n", c)

	var d bool
	d = true
	fmt.Printf("d=%v	\n", d)
	fmt.Printf("d=%+v	\n", d)
	fmt.Printf("d=%#v	\n", d)
}
