package main

import (
	"fmt"
	"os"

	"github.com/bioform/go-test/internal/build"
	"github.com/bioform/go-test/internal/greeting"

	"rsc.io/quote"
)

// Version of application
var Version = "development"

func main() {
	fmt.Println(quote.Go())
	fmt.Printf("Version %s\n", Version)
	fmt.Println("build.Time:\t", build.Time)
	fmt.Println("build.User:\t", build.User)

	name := "unknown"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	message, err := greeting.Greet(name)
	if err != nil {
		// fmt.Println(err.Error())
		// os.Exit(1)
		panic(err)
	}

	fmt.Println(message)
}
