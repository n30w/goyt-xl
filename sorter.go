package main

import (
	"fmt"
	"strconv"
)

// This file will sort the dates, given a year

type SortedVideos struct {
	List       []string
	ingestList []RelevantData
	Year       string
	YearCount  int
	ColumnLen  float64
}

// Exports data for use in xl.go
func (sv *SortedVideos) export() ([]RelevantData, string) {
	return sv.ingestList, sv.Year
}

// gets video field based on iteration via uint8 value
// This is probably a terrible way to do this. Switches aren't meant for this, I think. Lol.
func (sv *SortedVideos) getVideoField(index int, rv RelevantData) string {
	var field string
	switch index {
	case 0:
		field = rv.T
	case 1:
		field = rv.VC
	case 2:
		field = rv.PD
	}
	return field
}

// This function will ingest videos from  the excel sheet
func (sv *SortedVideos) ingestVideos() {
	data := readSpreadsheet()
	// rd := make([]RelevantData, len(data[0]))
	sv.Year = "2021"

	rd := make([]RelevantData, len(data[0]))
	// Starts at 1 since the very first index is the column header
	// removing https://youtube.com/... and just using the videoID
	for i := 1; i < len(data[1]); i++ {
		data[1][i] = data[1][i][32:]
	}

	data[0] = data[0][1:]
	data[1] = data[1][1:]

	// Get YouTube Data
	for i := 0; i < len(data[0]); i++ {
		d := RelevantData{
			T: data[0][i],
		}
		d.returnRelevantData(data[1][i])
		fmt.Println("PROCESSING:", data[0][i])
		rd[i] = d
	}

	sv.ingestList = rd
	// fmt.Println(sv.ingestList)
}

// Only look for target year
func (sv *SortedVideos) removeIrrelevantYears() {
	const COLUMNLENFACTOR = 0.83
	var sorted []RelevantData

	for i := 0; i < len(sv.ingestList); i++ {
		if sv.ingestList[i].PD == sv.Year {
			sorted = append(sorted, sv.ingestList[i])

			// Add to total year view count
			k, err := strconv.Atoi(sv.ingestList[i].VC)
			Enil(err)
			sv.YearCount += k

			// update length of column width
			if i != 0 {
				prev := len(sv.ingestList[i-1].T)
				cur := len(sv.ingestList[i].T)
				if cur > prev {
					sv.ColumnLen = float64(cur) * COLUMNLENFACTOR
				}
			} else {
				sv.ColumnLen = float64(len(sv.ingestList[0].T))
			}
		}
	}
	sv.ingestList = sorted
}
