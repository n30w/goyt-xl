package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type SearchedData struct {
	Year    string
	ListLen int
}

func (sd *SearchedData) SearchForThisYear(Year string) {
	sd.Year = Year
	sd.writeSpreadsheet()
}

func ReadSpreadsheet() [][]string {

	// Create a new slice, then add column URL to slices

	// Multi-dim array, title and url.
	spreadsheetData := make([][]string, 2)

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
	for i, col := range cols {
		for j, rowCell := range col {
			// Add title and URL to URL and title field.
			spreadsheetData[i][j] = rowCell[32:]
		}
	}
	return spreadsheetData
}

func (sd *SearchedData) writeSpreadsheet() {
	f := excelize.NewFile()
	index := f.NewSheet(fmt.Sprintf("%d", sd.Year))
	// f.SetCellValue()
	f.SetActiveSheet(index)

	if err := f.SaveAs("output.xlsx"); err != nil {
		fmt.Println(err)
	}
}

// Export the sorted list of data to new excel file
func Export(sv *SortedVideos) {

	func() {
		f := excelize.NewFile()
		index := f.NewSheet(sv.Year)
		f.SetActiveSheet(index)
		err := f.SaveAs("output.xlsx")
		Enil(err)
	}()
}
