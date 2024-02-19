package main

type BadX struct {
	// external module struct
}

func (x BadX) hello() string {
	return "bad"
}

type GoodX struct {
	// external module struct
}

func (x GoodX) hello() string {
	return "good"
}