package main

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type testGoodX1 struct{}
type testGoodX2 struct{}
type testGoodXDB struct{
	db *sql.DB
}

func (x testGoodX1) hello() string {
	return "This is test"
}

func (x testGoodX2) hello() string {
	return "This is another test"
}

func (x testGoodXDB) hello() string {
	return "This is db"
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

func TestHelloFromDB(t *testing.T) {
	db, err := sql.Open("sqlite3", "./example.sql")
	if err != nil {
		log.Fatal("failed to connect db: ", err)
	}
	defer db.Close()

	var (
		given IGoodX
		got   string
		want  string
	)
	given = testGoodXDB{db: db}
	got = helloFromGoodX(given)
	want = "This is db world"
	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}
