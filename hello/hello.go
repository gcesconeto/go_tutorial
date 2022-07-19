package main

import (
	"fmt"

	quote "rsc.io/quote"

	"github.com/gcesconeto/go_tutorial/greetings"
)

func main() {
	fmt.Println("Hello, Gopher!")
	fmt.Println(quote.Go())
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
