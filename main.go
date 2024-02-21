package main

import (
	"fmt"
)

func main() {
	// bad example
	fmt.Println(helloFromBadX())

	// good example
	var x IExt = Ext{}
	fmt.Println(helloFromGoodX(x))
}
