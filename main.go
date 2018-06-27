package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/yamatsuka-hiroto/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the monkey programming language!\n", user.Name)

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
