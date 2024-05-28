// Package main is the entry point of the Monkey programming language.
package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/w40141/monkey-language/golang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
