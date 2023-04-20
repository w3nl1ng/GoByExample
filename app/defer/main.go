package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("../../res/defer.txt")
	writeFile(f, []byte("defer"))
	closeFile(f)
}

func createFile(name string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, data []byte) int {
	fmt.Println("Writing")
	n, err := f.Write(data)
	if err != nil {
		panic(err)
	}
	return n
}

func closeFile(f *os.File) {
	fmt.Println("Closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}