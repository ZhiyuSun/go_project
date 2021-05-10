package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Go 语言的错误处理的方式，本质上是返回值检查，但是它也兼顾了异常的一些好处——对错误的扩展。

// 长度不够，少一个Weight
var b = []byte {0x48, 0x61, 0x6f, 0x20, 0x43, 0x68, 0x65, 0x6e, 0x00, 0x00, 0x2c}
var r = bytes.NewReader(b)

type Person struct {
	Name [10]byte
	Age uint8
	Weight uint8
	err error
}
func (p *Person) read(data interface{}) {
	if p.err == nil {
		p.err = binary.Read(r, binary.BigEndian, data)
	}
}

func (p *Person) ReadName() *Person {
	p.read(&p.Name)
	return p
}
func (p *Person) ReadAge() *Person {
	p.read(&p.Age)
	return p
}
func (p *Person) ReadWeight() *Person {
	p.read(&p.Weight)
	return p
}
func (p *Person) Print() *Person {
	if p.err == nil {
		fmt.Printf("Name=%s, Age=%d, Weight=%d\n",p.Name, p.Age, p.Weight)
	}
	return p
}

func main() {
	p := Person{}
	p.ReadName().ReadAge().ReadWeight().Print()
	fmt.Println(p.err)  // EOF 错误
}