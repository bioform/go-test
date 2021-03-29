package main

import (
	"github.com/bioform/go-test/internal/gameone"
	"github.com/bioform/go-test/internal/gameone/storage"
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

	firstScene := &book.Scene

	sceneId, err := storage.GetInstance().ReadLastPosition()
	if err != nil {
		panic(err)
	}
	if sceneId != nil {
		firstScene = book.Scenes[*sceneId]
	}

	err = book.Run(firstScene)
	if err != nil {
		panic(err)
	}
}
