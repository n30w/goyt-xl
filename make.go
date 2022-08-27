package main

func operate() error {
	sv := new(SortedVideos)
	sv.ingestVideos()
	sv.removeIrrelevantYears()
	writeSpreadsheet(sv)
	return nil
}
