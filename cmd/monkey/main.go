package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/uga-rosa/monkey/internal/repl"
)

func main() {
	user, err := user.Current()
	check(err)
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
