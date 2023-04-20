package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}


func main() {
	p := point{1, 2}

	var pf = fmt.Printf

	pf("struct1: %v\n", p)
	pf("struct2: %+v\n", p)
	pf("struct3: %#v\n", p)

	pf("type: %T\n", p)
	pf("bool: %t\n", true)
	pf("int: %d\n", 123)

	pf("bin: %b\n", 14)

	pf("char: %c\n", 33)

	pf("hex: %x\n", 456)
	pf("float: %f\n", 78.9)

	pf("float2: %e\n", 123400000.0)
	pf("float3: %E\n", 123400000.0)

	pf("str1: %s\n", "\"string\"")
	pf("str2: %q\n", "\"string\"")
	pf("str3: %x\n", "hex this")

	pf("pointer: %p\n", &p)
	pf("withd1: |%6d|%6d|\n", 12, 456)

	pf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	pf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	pf("width4: |%6s|%6s|\n", "foo", "b")
	pf("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf return a %s", "format string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}