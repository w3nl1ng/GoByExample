package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	d := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("d:", d)

	fmt.Println(filepath.Join("dir1//", "hello"))
	fmt.Println(filepath.Join("dir1/../dir1", "hello"))

	fmt.Println("dir:", filepath.Dir(d))
	fmt.Println("base:", filepath.Base(d))

	fmt.Println(filepath.IsAbs("dir1/hello"))
	fmt.Println(filepath.IsAbs("/dir1/hello"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel1, err := filepath.Rel("a/b", "a/b/dir1/filename")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel1)

	rel2, err := filepath.Rel("a/b", "a/c/dir1/filename")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel2)
}