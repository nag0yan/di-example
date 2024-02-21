package main

import "fmt"

func helloFromGoodX(x IExt) string {
	return fmt.Sprintf("%v world", x.hello())
}
