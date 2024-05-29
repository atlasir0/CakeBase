package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	ext := flag.String("ext", "", "file extension to search for")
	flag.Parse()

	directory := flag.Arg(0)
	if directory == "" {
		fmt.Println("Usage: ./myFind <directory>")
		os.Exit(1)
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if *ext == "" || filepath.Ext(path) == *ext {
			fmt.Println(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
