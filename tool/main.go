package main

import (
	"fmt"
	"log"
	"os"
)

type command func([]string) error

var commands = map[string]command{
	"hello": hello,
}

func hello(args []string) error {
	fmt.Println("Hello, World!")
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least one argument")
		return
	}

	cmd, ok := commands[os.Args[1]]
	if !ok {
		log.Fatal("Unknown command:", os.Args[1])
		return
	}

	if err := cmd(os.Args[2:]); err != nil {
		log.Fatal("Error:", err)
	}
}
