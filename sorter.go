package main

// Sorts dates given appropriate year and puts them into respective slices

import (
	"fmt"
	"strconv"
)

type SortedVideos struct {
	List         []string
	IngestVideos []RelevantData
	Year         string
	YearCount    int
	ColumnLen    float64
}

// Exports data for use in xl.go
func (sv *SortedVideos) export() ([]RelevantData, string) {
	return sv.IngestVideos, sv.Year
}

// Gets video field based on iteration via int value. This is probably a terrible way to do this. Switches aren't meant for this, I think. Lol.
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

// Ingests videos and puts them in appropriate slice
func (sv *SortedVideos) ingestVideos() {
	data := readSpreadsheet()
	rd := make([]RelevantData, len(data[0]))
	sv.Year = "2021"

	// Starts at 1 since the very first index is the column header
	for i := 1; i < len(data[1]); i++ {
		data[1][i] = data[1][i][32:] // Removes youtube.com url
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
	sv.IngestVideos = rd
}

// Remove everything but target year from slice
func (sv *SortedVideos) removeIrrelevantYears() {
	const COLUMNLENFACTOR = 0.83
	var sorted []RelevantData

	for i := 0; i < len(sv.IngestVideos); i++ {
		if sv.IngestVideos[i].PD == sv.Year {
			sorted = append(sorted, sv.IngestVideos[i])

			// Add to total year view count
			k, err := strconv.Atoi(sv.IngestVideos[i].VC)
			Enil(err)
			sv.YearCount += k

			// Update length of column width
			if i != 0 {
				prev := len(sv.IngestVideos[i-1].T)
				cur := len(sv.IngestVideos[i].T)
				if cur > prev {
					sv.ColumnLen = float64(cur) * COLUMNLENFACTOR
				}
			} else {
				sv.ColumnLen = float64(len(sv.IngestVideos[0].T))
			}
		}
	}
	sv.IngestVideos = sorted
}
