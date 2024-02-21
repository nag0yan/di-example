package main

type Ext struct {
	// external module struct
}

func (x Ext) hello() string {
	return "hello"
}

type IExt interface {
	hello() string
}
