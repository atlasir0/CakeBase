package comparedb

import (
	"bufio"
	"fmt"
	"os"
)

func readFileToMap(filename string) (map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	filesMap := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filesMap[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return filesMap, nil
}

func CompareFS(oldFilename, newFilename string) error {
	oldFiles, err := readFileToMap(oldFilename)
	if err != nil {
		return fmt.Errorf("error reading old snapshot: %w", err)
	}

	newFile, err := os.Open(newFilename)
	if err != nil {
		return fmt.Errorf("error reading new snapshot: %w", err)
	}
	defer newFile.Close()

	newFiles := make(map[string]bool)
	scanner := bufio.NewScanner(newFile)
	for scanner.Scan() {
		newFile := scanner.Text()
		newFiles[newFile] = true

		if _, exists := oldFiles[newFile]; !exists {
			fmt.Printf("ADDED %s\n", newFile)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning new snapshot: %w", err)
	}

	for oldFile := range oldFiles {
		if _, exists := newFiles[oldFile]; !exists {
			fmt.Printf("REMOVED %s\n", oldFile)
		}
	}

	return nil
}
