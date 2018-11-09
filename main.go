package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Interpreter/repl"
)

func main() {

	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Initializing command prompt for %s\n", currentUser.Username)
	repl.Start(os.Stdin, os.Stdout)
}
