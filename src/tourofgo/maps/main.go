package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var m1 = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m1)

	m2 := make(map[string]int)

	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])

	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])

	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	s := "This is a test test is a way to test functionality in a program"

	fmt.Println(WordCount(s))
}

func WordCount(s string) map[string]int {

	wcm := make(map[string]int)
	sf := strings.Fields(s)
	for _, word := range sf {
		wcm[word]++
	}
	return wcm
	//return map[string]int{"x": 1}
}
