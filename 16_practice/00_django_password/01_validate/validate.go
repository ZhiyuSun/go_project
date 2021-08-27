package main

import (
	"fmt"
	"github.com/alexandrevicenzi/unchained"
)

func main() {
	valid, err := unchained.CheckPassword("12345678", "pbkdf2_sha256$36000$iWOF4uGDZJcS$6SwEOwwjWq4z1PIw38D3hHX1j99sGoIRO2rfVwDxylQ=")

	if valid {
		fmt.Println("Password is valid.")
	} else {
		if err == nil {
			fmt.Println("Password is invalid.")
		} else {
			fmt.Printf("Error decoding password: %s\n", err)
		}
	}
}