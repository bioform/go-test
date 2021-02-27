package main

import (
	"github.com/bioform/go-test/internal/gameone"
	"github.com/inancgumus/screen"
)

// Version of application
var Version = "development"

func main() {
	screen.Clear()
	screen.MoveTopLeft()

	book, err := gameone.LoadBook("book.yml")
	if err != nil {
		panic(err)
	}

	book.Run()
}
