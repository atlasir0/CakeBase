package pkg

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func CountWords(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка открытия файла %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	fmt.Printf("%d\t%s\n", wordCount, filename)
}

func CountLines(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка открытия файла %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	fmt.Printf("%d\t%s\n", lineCount, filename)
}

func CountChars(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка открытия файла %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения информации о файле %s: %v\n", filename, err)
		return
	}

	fmt.Printf("%d\t%s\n", fileInfo.Size(), filename)
}
