package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, args ...string) error {
	// fmt.Println("Thank you for visiting and have a wonderful day!")
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
