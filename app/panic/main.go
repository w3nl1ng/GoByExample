package main

import "os"

func main() {
	panic("a problem")

	// 地道的做法 in go
	_, err := os.Create("/temp/file")
	if err != nil {
		panic(err)
	}
}