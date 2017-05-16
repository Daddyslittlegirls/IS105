package main

import "fmt"
import "os"
import "strconv"
import "./log"
//import "./log.go"


// Tar 2 argumenter (float & int) og kjører Logbcli funksjonen med de.
func main() {
  arg1 := os.Args[1]
	arg2 := os.Args[2]

	argFloat, err := strconv.ParseFloat(arg1, 64)
	argInt, err := strconv.ParseInt(arg2, 10, 64)

	if err == nil {
		log.Logbcli(argFloat, argInt)
	}
}
