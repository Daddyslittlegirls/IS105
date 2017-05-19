package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Skriv jp for japansk oversettelse eller is for islandsk oversettelse", "\n")
	fmt.Print("\xE2\x80\x9C\x6E\x6F\x72\x64\x20\x6F\x67\x20\x73\xC3\xB8\x72\xE2\x80\x9D", "\n")
	for scanner.Scan() {
	if scanner.Text() == "jp" {
			fmt.Print("på japansk er ")
			fmt.Print("\xE2\x80\x9C\xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97\xE2\x80\x9D", "\n")
		} else {
			if scanner.Text() == "is" {
			fmt.Print("på islandsk er ")
			fmt.Print("\xE2\x80\x9C\x6E\x6F\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\xE2\x80\x9D", "\n")
			}
		}
	}
}
