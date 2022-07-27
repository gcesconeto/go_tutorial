package main

import (
	"fmt"
	"log"

	"github.com/gcesconeto/go_tutorial/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, Gopher!")
	fmt.Println(quote.Go())

	log.SetPrefix("greetings error: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Gopher")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
