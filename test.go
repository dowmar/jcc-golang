package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	data := []byte(`{"x": 1, "y": 2}`)
	var p Point
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(p)
	}
}
