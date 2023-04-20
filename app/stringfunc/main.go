package main

import (
	"fmt"
	s "strings"
)

func main() {
	var p = fmt.Println

	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("test", "t", "0", -1))
	p("Replace:", s.Replace("test", "t", "0", 1))
	p("Split:", s.Split("t-e-s-t", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper", s.ToUpper("test"))
}