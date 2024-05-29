package main

import (
	"flag"
	"fmt"
	"myWc/pkg"
	"os"
	"sync"
)

func main() {

	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
	countChars := flag.Bool("m", false, "Count characters")
	flag.Parse()

	if !*countLines && !*countWords && !*countChars {
		*countWords = true
	}

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Не указаны файлы для обработки")
		os.Exit(1)
	}

	var countFunc func(string, *sync.WaitGroup)
	switch {
	case *countLines:
		countFunc = pkg.CountLines
	case *countChars:
		countFunc = pkg.CountChars
	default:
		countFunc = pkg.CountWords
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go countFunc(file, &wg)
	}

	wg.Wait()
}
