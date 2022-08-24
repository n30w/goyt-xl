package main

// This file will sort the dates, given a year

type SortedVideos struct {
	List       []string
	ingestList []RelevantData
	Year       string
}

// This function will ingest videos from  the excel sheet
func (sv *SortedVideos) ingestVideos() {
	rd := []RelevantData{}
	data := ReadSpreadsheet()

	// Populates slice rd with RelevantData structs
	for i, url := range data {
		d := RelevantData{
			T: data[i][0],
		}
		d.ReturnRelevantData(url[1])
		rd[i] = d
	}
	sv.ingestList = rd
}

// Remove irrelevant years
func (sv *SortedVideos) removeYears() {
	sorted := []RelevantData{}
	for _, v := range sv.ingestList {
		if v.PD[:3] == sv.Year {
			sorted = append(sorted, v)
		}
	}
	sv.ingestList = sorted
}

// TODO:
// Export sorted data to excel sheet
// Write to excel sheet
