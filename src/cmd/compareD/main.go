package main

import (
	compare "Go_Day01/cmd/compareD/compare"
	reader "Go_Day01/cmd/readDB/reader"
	"flag"
	"fmt"
)

func main() {
	oldFilename := flag.String("old", "", "old database file to compare")
	newFilename := flag.String("new", "", "new database file to compare")
	flag.Parse()

	if *oldFilename == "" || *newFilename == "" {
		fmt.Println("Please provide both old and new database files")
		return
	}

	oldRecipe, err := reader.ReadDB(*oldFilename)
	if err != nil {
		fmt.Println("Error reading old database file:", err)
		return
	}

	newRecipe, err := reader.ReadDB(*newFilename)
	if err != nil {
		fmt.Println("Error reading new database file:", err)
		return
	}

	compare.CompareCakes(oldRecipe, newRecipe)
}
