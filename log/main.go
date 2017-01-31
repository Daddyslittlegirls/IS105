package main

import "fmt"
import "os"
import "strconv"
//import "./log.go"

func main() {
  log2(20)

  arg1 := os.Args[1]
	arg2 := os.Args[2]
  fmt.Println(arg1, arg2)
	
	argFloat, err := strconv.ParseFloat(arg1, 64)
	argInt, err := strconv.ParseInt(arg2, 10, 64)
  
	if err == nil {
		logbcli(argFloat, argInt)
	}
}
