package main

import "fmt"
import "os"
import "math"
import "strconv"
//import "./log.go"

func calculateLog(x float64){
	fmt.Println(math.Log2(x))
}

func main() {

    arg := os.Args[1]

	argFloat, err := strconv.ParseFloat(arg, 64)
	if err == nil {
		calculateLog(argFloat)
	}
}
