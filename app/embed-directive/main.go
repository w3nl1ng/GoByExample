package main

import (
	"embed"
	"fmt"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var res embed.FS

func main() {
	fmt.Println(fileString)
	fmt.Println(string(fileByte))

	content1, _ := res.ReadFile(`folder/file1.hash`)
	fmt.Println(string(content1))

	content2, _ := res.ReadFile(`folder/file2.hash`)
	fmt.Println(string(content2))
}
