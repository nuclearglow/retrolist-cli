package main

import (
	"fmt"
	"math/rand"

	"retrocli.svenvowe.de/config"
	"retrocli.svenvowe.de/retrolist"
)

func main() {
	fmt.Printf("Welcome to Retro CLI\n")

	// create some test data for now
	list := retrolist.NewRetroList("Test List", "Testing")

	for i := range 10 {
		item := retrolist.NewItem(fmt.Sprintf("Item %d", i+1), uint(rand.Intn(i+1)+1))
		list.AddItem(item)
	}

	err := list.Save(config.DefaultFilename)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
