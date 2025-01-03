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
	// paths = []lemin.Path{
	// 	{Path: []string{"1", "d", "3"}},
	// 	{Path: []string{"1", "5", "3"}},
	// 	{Path: []string{"1", "8", "5", "3"}},
	// 	{Path: []string{"1", "7", "9", "3"}},
	// 	{Path: []string{"1", "a", "h", "3"}},
	// }
	filterpath := lemin.FilterPaths(paths, lemin.Ants)

	var result [][]string
	// Step 2: Store the Values from each Path struct into the 2D slice
	for _, path := range filterpath {
		result = append(result, path.Path)
	}

	lemin.Sumilation(lemin.Ants, result)
}
