package main

func operate() error {
	sv := new(SortedVideos)
	sv.ingestVideos()
	sv.removeYears()
	writeSpreadsheet(sv)
	return nil
}
