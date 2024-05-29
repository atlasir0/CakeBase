package finder

import (
	"fmt"
	"myFind/config"
	"os"
	"path/filepath"
)

func TraverseDirectory(params config.Parameters, flags config.Flags) {
	filepath.Walk(params.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if flags.ShowSymlinks && isSymlink(info) {
			printEntry(path, info)
		} else if flags.ShowDirs && info.IsDir() {
			printEntry(path, info)
		} else if flags.SpecificExt && filepath.Ext(path) == ("."+params.Extension) {
			printEntry(path, info)
		} else if flags.ShowFiles && !flags.SpecificExt && isFile(info) {
			printEntry(path, info)
		} else if noFlags(flags) {
			printEntry(path, info)
		}
		return nil
	})
}
func isFile(info os.FileInfo) bool {
	return info.Mode().IsRegular()
}

func isSymlink(info os.FileInfo) bool {
	return info.Mode()&os.ModeSymlink == os.ModeSymlink
}

func noFlags(flags config.Flags) bool {
	return !flags.ShowDirs && !flags.ShowFiles && !flags.SpecificExt && !flags.ShowSymlinks
}

func printEntry(path string, info os.FileInfo) {
	if isSymlink(info) {
		targetPath, err := os.Readlink(path)
		if err != nil {
			targetPath = "[broken]"
		}
		fmt.Println(path, "->", targetPath)
	} else {
		fmt.Println(path)
	}
}
