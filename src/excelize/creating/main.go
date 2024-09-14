package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	file := excelize.NewFile()

	headers := []string{"ID", "Name", "Age"}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	data := [][]interface{}{
		{1, "John", 30},
		{2, "Alex", 20},
		{3, "Bob", 40},
	}

	for i, row := range data {
		dataRow := i + 2
		for j, col := range row {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	if err := file.SaveAs("students.xlsx"); err != nil {
		log.Fatal(err)
	}
}
