package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name not allowed")
	}

	message := fmt.Sprintf(randomMessage(), name)

	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomMessage() string {
	messages := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return messages[rand.Intn(len(messages))]
}
