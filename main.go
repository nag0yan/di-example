package main

import (
	"fmt"
)

func main() {
	// bad example
	fmt.Println(helloFromBadX())

	// good example
	var x IGoodX = Ext{}
	fmt.Println(helloFromGoodX(x))
}
