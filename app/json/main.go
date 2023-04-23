package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type resp1 struct {
	Page int
	Fruits []string
}

type resp2 struct {
	Page int        `json:"page"`
	Fruits []string  `json:"fruits"`
}

func main() {
	bolJ, _ := json.Marshal(true)
	fmt.Println(string(bolJ))

	intJ, _ := json.Marshal(1)
	fmt.Println(string(intJ))

	strJ, _ := json.Marshal("hello")
	fmt.Println(string(strJ))

	sliJ, _ := json.Marshal([]string{"hello", "world", "jack"})
	fmt.Println(string(sliJ))

	mapJ, _ := json.Marshal(map[string]int{"jack": 1, "rose": 2, "tom": 3})
	fmt.Println(string(mapJ))

	resp1D := resp1{
		Page: 1,
		Fruits: []string{"apple", "banana", "peach"},
	}
	resp1J, _ := json.Marshal(resp1D)
	fmt.Println(string(resp1J))

	resp2D := resp2{
		Page: 2,
		Fruits: []string{"apple", "pear", "banana"},
	}
	resp2J, _ := json.Marshal(resp2D)
	fmt.Println(string(resp2J))

	byt := []byte(`{"num": 6.12, "strs":["a", "b"]}`)
	var data map[string]interface{}
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	num := data["num"].(float64)
	fmt.Println(num)

	strs := data["strs"].([]interface{})
	str0 := strs[0].(string)
	fmt.Println(str0)

	resp1Data := []byte(`{"Page": 1, "fruits":["apple", "banana"]}`)
	resp1J1 := resp1{}
	if err := json.Unmarshal(resp1Data, &resp1J1); err != nil {
		panic(err)
	}
	fmt.Println(resp1J1)
	fmt.Println(resp1J1.Fruits[0])

	jn := json.NewEncoder(os.Stdout)
	out := map[string]string{"jack":"tom", "apple":"banana"}
	jn.Encode(out)
}