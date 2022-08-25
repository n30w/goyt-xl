package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func readSpreadsheet() [2][]string {
	// Multi-dim array, title and url.
	spreadsheetData := [2][]string{}

	f, err := excelize.OpenFile("apsyl.xlsx")
	Enil(err)
	defer func() {
		// Close the spreadsheet
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	cols, err := f.GetCols("punahou")
	Enil(err)

	// Goes by each column and returns row value
	for i := range cols[:2] {
		for j := 1; j < len(cols[i:2])-1; j++ {
			// Add title and URL to title and URL field.
			if i == 1 {
				spreadsheetData[i][j] = cols[i][j][32:]
			} else {
				spreadsheetData[i][j] = cols[i][j]
			}
		}
	}
	return spreadsheetData
}

func writeSpreadsheet(sv *SortedVideos) {
	rd, year := sv.export()
	columnAxis := [3]string{
		"A",
		"B",
		"C",
	}

	f := excelize.NewFile()
	index := f.NewSheet(year)
	f.SetActiveSheet(index)

	// Can use go routines here to do multiple column work.
	// Writes out all sorted data to excel rows
	for i := 0; i < len(columnAxis); i++ {
		for j := range rd {
			axis := columnAxis[i] + strconv.Itoa(j)
			err := f.SetCellStr(sv.Year, axis, sv.getVideoField(j, rd[i]))
			Enil(err)
		}
	}

	if err := f.SaveAs("output.xlsx"); err != nil {
		fmt.Println(err)
	}
}
