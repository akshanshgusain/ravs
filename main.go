package main

import (
	"fmt"
	"os"
	"ravs_lang/repl"
)

func main() {
	fmt.Printf("Hi! I am Ravs\n")
	repl.Start(os.Stdin, os.Stdout)
}
