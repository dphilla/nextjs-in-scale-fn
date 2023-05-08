package main

import (
    "embed"
    "fmt"
)


//go:embed *
var folder embed.FS

func main() {
	fileContents, _ := folder.ReadFile("404.html")
	fmt.Println(fileContents)
}
