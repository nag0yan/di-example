package main

import "fmt"

type IGoodX interface {
	hello() string
}

func helloFromGoodX(x IGoodX) string {
	return fmt.Sprintf("%v world", x.hello())
}
