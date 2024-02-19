package main

import "fmt"

func helloFromBadX() string {
	x := BadX{}
	return fmt.Sprintf("%v world", x.hello())
}
