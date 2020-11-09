package greeting

import (
	"errors"
	"fmt"
)

// ErrEmptyName - empty "name" attribute is restricted
var ErrEmptyName = errors.New("NAME cannot be empty")

// Greet - display hellow string
func Greet(name string) (string, error) {
	if len(name) == 0 {
		return "", ErrEmptyName
	}

	return compileGreeting(name), nil
}

func compileGreeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
