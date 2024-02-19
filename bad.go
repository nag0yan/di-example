package main

import "fmt"

func helloFromBadX() string {
	x := Ext{}
	return fmt.Sprintf("%v world", x.hello())
}
