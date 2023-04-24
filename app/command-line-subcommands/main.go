package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooStr := fooCmd.String("str", "", "foo string")
	fooEnable := fooCmd.Bool("enable", false, "foo enable")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barNum := barCmd.Int("num", 0, "bar num")

	if len(os.Args) < 2 {
		fmt.Println("need a subcommand(foo or bar)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println(*fooStr)
		fmt.Println(*fooEnable)
		fmt.Println(fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println(*barNum)
		fmt.Println(barCmd.Args())
	default:
		fmt.Println("subcommand should be foo or bar")
		os.Exit(1)
	}

}
