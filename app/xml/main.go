package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id int           `xml:"id,attr"`
	Name string      `xml:"name"`
	Origin []string    `xml:"origin"`
}


func (p Plant)String() string {
	return fmt.Sprintf("id=%v, name=%v, origin=%v",
						p.Id, p.Name, p.Origin)
}

func main() {
	coffee := Plant{Id: 27, Name: "coffee"}
	coffee.Origin = []string{"apple", "banana"}
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))
}