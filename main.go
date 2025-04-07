package main

import (
	"fmt"
	"math/rand"

	"retrocli.svenvowe.de/config"
	"retrocli.svenvowe.de/retrolist"
)

func main() {
	fmt.Printf("Welcome to Retro 2 CLI\n")

	// create some test data for now
	list := retrolist.NewRetroList("Test List 2", "Testing 2")

	for i := range 10 {
		item := retrolist.NewItem(fmt.Sprintf("Item %d", i+1), uint(rand.Intn(i+1)+1))
		list.AddItem(item)
	}

	fmt.Printf("Saving RetroList '%s'\n", list.Title)
	err := list.Save(config.DefaultFilename)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
