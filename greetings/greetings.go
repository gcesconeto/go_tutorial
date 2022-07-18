package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Ohai, %v. Welcome!", name)

	return message
}
