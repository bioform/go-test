package greeting

import (
	"fmt"

	"github.com/go-errors/errors"
)

// ErrEmptyName - empty "name" attribute is restricted
var ErrEmptyName = errors.Errorf("NAME cannot be empty")

// Greet - display hellow string
func Greet(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.Wrap(ErrEmptyName, 1)
	}

	return compileGreeting(name), nil
}

func compileGreeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
