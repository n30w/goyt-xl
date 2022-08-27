package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func readSpreadsheet() [][]string {
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
	// fmt.Println(len(cols[0]))
	// fmt.Println(cols[1][3][32:])
	return cols[:2]
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

	// Dude what is this
	err := f.SetCellStr(year, "A1", "TITLE")
	Enil(err)
	err = f.SetCellStr(year, "B1", "VIEWS")
	Enil(err)
	err = f.SetCellStr(year, "C1", "YEAR")
	Enil(err)
	err = f.SetCellStr(year, "D1", "TOTAL")
	Enil(err)

	// Output total year count
	err = f.SetCellValue(year, "D2", sv.YearCount)
	Enil(err)

	// Can use go routines here to do multiple column work.

	// Writes out all sorted data to excel rows

	// Adds titles to A column
	for i := 0; i < 1; i++ {
		for j := range rd {
			axis := columnAxis[i] + strconv.Itoa(j+2)
			err := f.SetCellStr(year, axis, sv.getVideoField(i, rd[j]))
			Enil(err)
		}
	}
	// Convert strings to int
	for i := 1; i < 3; i++ {
		for j := range rd {
			axis := columnAxis[i] + strconv.Itoa(j+2)
			k, err := strconv.Atoi(sv.getVideoField(i, rd[j]))
			Enil(err)
			err = f.SetCellValue(year, axis, k)
			Enil(err)
		}
	}

	// Set column width of title column
	err = f.SetColWidth(year, "A", "A", sv.ColumnLen)
	Enil(err)

	if err := f.SaveAs("output.xlsx"); err != nil {
		fmt.Println(err)
	}
}
