package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("you")
	fmt.Println(message)
}
