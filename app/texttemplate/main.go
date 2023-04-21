package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("values is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// t1 = template.Must(t1.Parse("value is {{.}}\n"))

	t1.Execute(os.Stdout, "some data")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"c++",
		"rust",
		"go",
		"java",
	})

	create := func(name, temp string) *template.Template {
		return template.Must(template.New(name).Parse(temp))
	}

	t2 := create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{Name: "tom jack"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "hello world",
	})

	t3 := create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "mot empty")
	t3.Execute(os.Stdout, "")

	t4 := create("t4", "{{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"hello",
		"jack",
		"tom",
	})
}