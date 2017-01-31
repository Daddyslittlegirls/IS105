package main

import "fmt"
import "os"
import "math"
//import "./log.go"

func main() {
    arg := os.Args[1]
    calculateLog(arg);
}

func calculateLog(x float64){
	fmt.Println(math.Log2(x))
}
