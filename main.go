package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/tayyabtalha/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s Welcome to interpretor\n", user)
	fmt.Println("Developed by Tayyab Talha mtayyabtalha@gmail.com")
	repl.Start(os.Stdin, os.Stdout)
}
