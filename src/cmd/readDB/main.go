package main

import (
	reader "Go_Day01/cmd/readDB/reader"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"strings"
)

func main() {
	filename := flag.String("f", "", "file to read")
	flag.Parse()
	reader.Check(*filename)

	recipe, err := reader.ReadDB(*filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if strings.HasSuffix(*filename, ".json") {
		output, _ := xml.MarshalIndent(recipe, "", "    ")
		fmt.Println(string(output))
	} else {
		output, _ := json.MarshalIndent(recipe, "", "    ")
		fmt.Println(string(output))
	}
}
