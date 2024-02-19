package main

import "testing"

type testGoodX1 struct{}
type testGoodX2 struct{}

func (x testGoodX1) hello() string {
	return "This is test"
}

func (x testGoodX2) hello() string {
	return "This is another test"
}

func TestHelloFromGoodX(t *testing.T) {
	var (
		given IGoodX
		got   string
		want  string
	)
	given = testGoodX1{}
	got = helloFromGoodX(given)
	want = "This is test world"
	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}

	given = testGoodX2{}
	got = helloFromGoodX(given)
	want = "This is another test world"
	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}
