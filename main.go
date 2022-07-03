package main

import (
	"fmt"
	"monke/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s, to the Monke programming language. Together we will return to monke.\n",
		user.Username)
	fmt.Printf("Begin typing in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
