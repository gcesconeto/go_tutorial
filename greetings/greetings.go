package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name not allowed")
	}

	message := fmt.Sprintf("Ohai, %v. Welcome!", name)

	return message, nil
}
