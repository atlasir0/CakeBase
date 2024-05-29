package main

import (
	"flag"
	"fmt"
	"log"
	"myRotate/config"
	"myRotate/pkg"
	"os"
)

func main() {
	archiveDir := flag.String("a", "", "directory to store archived log files")
	flag.Parse()

	if *archiveDir != "" {
		if _, err := os.Stat(*archiveDir); os.IsNotExist(err) {
			log.Fatalf("Error: archive directory %s does not exist", *archiveDir)
		}
	}

	logFiles := flag.Args()
	if len(logFiles) == 0 {
		fmt.Println("Usage: ./myRotate [-a archive_dir] log_file1 log_file2 ...")
		os.Exit(1)
	}

	config.LoadConfig()

	fileCh := make(chan string)

	for _, file := range logFiles {
		go pkg.RotateLog(file, *archiveDir, fileCh)
	}

	for range logFiles {
		fmt.Println(<-fileCh)
	}
}
