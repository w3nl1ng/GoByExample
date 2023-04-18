package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("../../res/in.txt")
	check(err)
	fmt.Println(string(data))

	f, err := os.Open("../../res/in.txt")
	check(err)

	b1 := make([]byte, 5)
	num, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", num, string(b1[:num]))

	pos1, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	num2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes st %d: %s\n", num2, pos1, string(b2[:num2]))

	pos2, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	num3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes at %d: %s\n", num3, pos2, string(b3[:num3]))

	_, err = f.Seek(0, 0)
	check(err)

	reader := bufio.NewReader(f)
	b4, err := reader.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	//Peek不会自动增加文件位置指针
	b4, err = reader.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	_, err = f.Seek(0, 0)
	check(err)

	//Read会自动增加文件位置指针
	b5 := make([]byte, 5)
	num4, err := reader.Read(b5)
	check(err)
	fmt.Printf("%d bytes: %s\n",num4, string(b5))

	num5, err := reader.Read(b5)
	check(err)
	fmt.Printf("%d bytes: %s\n",num5, string(b5))
}