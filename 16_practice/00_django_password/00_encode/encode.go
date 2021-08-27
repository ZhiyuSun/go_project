package main

import (
	"fmt"
	"github.com/alexandrevicenzi/unchained"
)

func main() {
	hash, err := unchained.MakePassword("12345678", unchained.GetRandomString(12), "default")

	if err == nil {
		fmt.Println(hash)
	} else {
		fmt.Printf("Error encoding password: %s\n", err)
	}
}