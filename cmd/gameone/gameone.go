package main

import (
	"github.com/bioform/go-test/internal/gameone"
)

// Version of application
var Version = "development"

func main() {
	scene, err := gameone.LoadScene("scene.yml")
	if err != nil {
		panic(err)
	}

	scene.Run()
}
