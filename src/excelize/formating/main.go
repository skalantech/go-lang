package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

// Style the file
func main() {

	file, err := excelize.OpenFile("students.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	style, err1 := file.NewStyle(
		&excelize.Style{
			Alignment: &excelize.Alignment{Horizontal: "center"},
			Font:      &excelize.Font{Bold: true, Color: "FF0000"},
			Border: []excelize.Border{
				{Type: "left", Color: "00FF0000", Style: 1},
				{Type: "right", Color: "00FF0000", Style: 1},
				{Type: "top", Color: "00FF0000", Style: 1},
				{Type: "bottom", Color: "00FF0000", Style: 1},
			},
			Fill: excelize.Fill{
				Type:    "pattern",
				Color:   []string{"#DDEBF7"},
				Pattern: 1,
			},
		},
	)
	if err1 != nil {
		log.Fatal(err1)
	}

	styleTitle, err2 := file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type: "pattern",
			// navy blue
			Color:   []string{"#000080"},
			Pattern: 1,
		},
		Font:      &excelize.Font{Bold: true, Color: "FFFFFF"},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	if err2 != nil {
		log.Fatal(err2)
	}

	file.SetCellStyle("Sheet1", "A1", "I5", style)
	file.SetCellStyle("Sheet1", "A1", "I1", styleTitle)
	// Adjust the range for applying style to chart
	file.SetCellStyle("Sheet1", "H5", "I5", styleTitle)

	if err4 := file.AddChart("Sheet1", "B8", &excelize.Chart{
		Type: excelize.Line,
		Series: []excelize.ChartSeries{
			{
				Name:       "Sheet1!$D$2",
				Categories: "Sheet1!$A$3:$A$5",
				Values:     "Sheet1!$D$3:$D$5",
			},
			{
				Name:       "Sheet1!$E$2",
				Categories: "Sheet1!$A$3:$A$5",
				Values:     "Sheet1!$E$3:$E$5",
			},
			{
				Name:       "Sheet1!$F$2",
				Categories: "Sheet1!$A$3:$A$5",
				Values:     "Sheet1!$F$3:$F$5",
			},
			{
				Name:       "Sheet1!$G$2",
				Categories: "Sheet1!$A$3:$A$5",
				Values:     "Sheet1!$G$3:$G$5",
			},
		},
	}); err4 != nil {
		log.Fatal(err4)
	}

	if err3 := file.Save(); err3 != nil {
		log.Fatal(err3)
	}
}
