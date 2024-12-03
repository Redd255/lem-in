package main

import (
	"fmt"
	lemin "lem-in/src"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Invalid arguments\nUsage : go run . <filename>")
	}

	colony := lemin.Colony{}
	if err := colony.ParseFile(args[0]); err != nil {
		fmt.Println("Error:", err)
		return
	}
	paths, err := colony.FindPaths()
	if err != nil {
        fmt.Println("Error:", err)
        return
    }
	for _, path := range paths {
		fmt.Println(path)
	}
}
