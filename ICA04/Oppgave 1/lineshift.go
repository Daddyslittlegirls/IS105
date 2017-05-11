package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("text1.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%q", b)
}
