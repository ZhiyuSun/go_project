package main

type Base struct {
	a string
}

type SonA struct {
	base Base
	name string
}

func main() {
	var list []interface{}
	s1 := SonA{
		base: Base{"1"},
		name: "s1",
	}
	list = append(list, s1)
}
