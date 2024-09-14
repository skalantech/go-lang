package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		if parseErr, ok := err.(*csv.ParseError); ok {
			log.Printf("CSV parse error at line %d, column %d: %v", parseErr.Line, parseErr.Column, parseErr.Err)
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println(records)
}
