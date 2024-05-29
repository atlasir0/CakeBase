package reader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Ingredient struct {
	Name  string `json:"ingredient_name,omitempty" xml:"itemname,omitempty"`
	Count string `json:"ingredient_count,omitempty" xml:"itemcount,omitempty"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

type Cake struct {
	Name        string       `json:"name,omitempty" xml:"name,omitempty"`
	Time        string       `json:"time,omitempty" xml:"stovetime,omitempty"`
	Ingredients []Ingredient `json:"ingredients,omitempty" xml:"ingredients>item,omitempty"`
}

type Recipe struct {
	Cakes []Cake `json:"cake,omitempty" xml:"cake,omitempty"`
}

func ReadDB(filename string) (*Recipe, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var recipe Recipe
	if strings.HasSuffix(filename, ".json") {
		json.Unmarshal(byteValue, &recipe)
	} else if strings.HasSuffix(filename, ".xml") {
		xml.Unmarshal(byteValue, &recipe)
	}
	return &recipe, nil
}

func Check(filename string) {
	if filename == "" {
		fmt.Println("Please provide a file name with -f flag")
		return
	}

	if !strings.HasSuffix(filename, ".json") && !strings.HasSuffix(filename, ".xml") {
		fmt.Println("Unsupported file type")
		return
	}
}
