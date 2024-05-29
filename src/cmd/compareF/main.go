package main

import (
	comparedb "Go_Day01/cmd/compareF/comparef"
	"flag"
	"fmt"
)

func main() {
	oldFilename := flag.String("old", "", "old filesystem snapshot to compare")
	newFilename := flag.String("new", "", "new filesystem snapshot to compare")
	flag.Parse()

	if *oldFilename == "" || *newFilename == "" {
		fmt.Println("Please provide both old and new filesystem snapshot files")
		return
	}

	if err := comparedb.CompareFS(*oldFilename, *newFilename); err != nil {
		fmt.Printf("Error comparing filesystems: %s\n", err)
	}
}
