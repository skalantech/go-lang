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
		log.Fatal(err)
	}

	fmt.Println(records)
}
