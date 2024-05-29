package pkg

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func RotateLog(logFile, archiveDir string, fileCh chan<- string) {
	fileInfo, err := os.Stat(logFile)
	if err != nil {
		fileCh <- fmt.Sprintf("Error: %v", err)
		return
	}

	timestamp := fileInfo.ModTime().Unix()
	archiveFilename := fmt.Sprintf("%s_%d.tar.gz", filepath.Base(logFile), timestamp)

	if archiveDir != "" {
		archiveFilename = filepath.Join(archiveDir, archiveFilename)
	} else {
		archiveFilename = filepath.Join(filepath.Dir(logFile), archiveFilename)
	}

	logFilePtr, err := os.Open(logFile)
	if err != nil {
		fileCh <- fmt.Sprintf("Error: %v", err)
		return
	}
	defer logFilePtr.Close()

	archiveFilePtr, err := os.Create(archiveFilename)
	if err != nil {
		fileCh <- fmt.Sprintf("Error: %v", err)
		return
	}
	defer archiveFilePtr.Close()

	gzipWriter := gzip.NewWriter(archiveFilePtr)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, logFilePtr)
	if err != nil {
		fileCh <- fmt.Sprintf("Error: %v", err)
		return
	}

	fileCh <- fmt.Sprintf("File %s archived successfully", logFile)
}
