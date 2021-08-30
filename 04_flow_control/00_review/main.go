package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func printFile1() {
	const filename = "04_flow_control/00_review/abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func printFile2() {
	const filename = "04_flow_control/00_review/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func grade(score int) string {
	g := ""
	switch  {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

// loop
func convertToBin(n int) string {
	result := ""
	for ; n>0; n/=2 {
		lsb := n%2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile3(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
// go语言经常要用到死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	printFile1()
	printFile2()
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		//grade(101),
		)
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		)
	printFile3("04_flow_control/00_review/abc.txt")

}
