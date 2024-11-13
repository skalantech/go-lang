package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Person struct {
	ID   int
	Name string
	Age  int
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var people []Person
	for _, record := range records[1:] {
		id, _ := strconv.Atoi(record[0])
		age, _ := strconv.Atoi(record[2])
		person := Person{
			ID:   id,
			Name: record[1],
			Age:  age,
		}
		people = append(people, person)
	}

	fmt.Println(people)
}
