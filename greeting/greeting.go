package greeting

import (
	"errors"
	"fmt"
)

// Greet - display hellow string
func Greet(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("NAME cannot be empty")
	}

	return compileGreeting(name), nil
}

func compileGreeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
