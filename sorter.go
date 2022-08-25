package main

// This file will sort the dates, given a year

type SortedVideos struct {
	List       []string
	ingestList []RelevantData
	Year       string
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
		field = rv.PD
	case 2:
		field = rv.VC
	}
	return field
}

// This function will ingest videos from  the excel sheet
func (sv *SortedVideos) ingestVideos() {
	rd := []RelevantData{}
	data := readSpreadsheet()

	// TODO:
	// Fix this for loop.
	// Returns error: "index out of range [0] with length 0"
	// Populates slice rd with RelevantData structs
	for i, url := range data {
		d := RelevantData{
			T: data[i][0],
		}
		d.returnRelevantData(url[1])
		rd[i] = d
	}
	sv.ingestList = rd
}

// Remove irrelevant years and return new array
func (sv *SortedVideos) removeYears() {
	sorted := []RelevantData{}
	for _, v := range sv.ingestList {
		if v.PD[:3] == sv.Year {
			sorted = append(sorted, v)
		}
	}
	sv.ingestList = sorted
}
