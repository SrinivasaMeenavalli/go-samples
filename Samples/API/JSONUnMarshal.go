package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	byt := []byte(`{"num":6.13,"strs":["pen","pencil"]}`)
	var data map[string]interface{}
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	num := data["num"].(float64)
	fmt.Println(num)

	strs := data["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
