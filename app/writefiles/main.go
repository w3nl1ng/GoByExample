package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	b1 := []byte("hello\ngo\n")
	err := os.WriteFile("../../res/out.txt", b1, 0664)
	check(err)

	f, err := os.Create("../../res/out2.txt")
	check(err)
	defer f.Close()

	data1 := []byte{115, 111, 109, 101, 10}
	n, err := f.Write(data1)
	check(err)
	fmt.Printf("writed %d bytes\n", n)

	n2, err := f.WriteString("hello")
	check(err)
	fmt.Printf("writed %d bytes\n", n2)

	writer := bufio.NewWriter(f)
	n3, err := writer.WriteString(" world\n")
	check(err)
	fmt.Printf("writed %d bytes\n", n3)

	writer.Flush()
}