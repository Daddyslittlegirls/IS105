package log

import "fmt"
import "math"

func Log2(x float64){
	fmt.Println(math.Log2(x))
}

func Logcli(x float64){
	fmt.Println(math.Log2(x))
}

func Logbcli(x float64, y int64) {
	if y == 2 {
		fmt.Println(math.Log2(x))
	} else if y == 10 {
		fmt.Println(math.Log10(x))
	} else {
		fmt.Println("Ugyldig nummer")
	}
}cd
