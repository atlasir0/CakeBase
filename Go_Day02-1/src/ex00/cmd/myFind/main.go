package main

import (
	"fmt"
	"myFind/config"
	finder "myFind/pkg"
)

func main() {
	params, flags, err := config.ParseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	finder.TraverseDirectory(params, flags)
}
