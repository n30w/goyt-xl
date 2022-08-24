/*
Retrieve YouTube data
*/

package main

import (
	"github.com/kkdai/youtube/v2"
)

// This is the only relevant data we need to make the excel sheet
type RelevantData struct {
	T  string // Title
	PD string // Publish Date
	VC string // View Count
}

// Returns RelevantData struct, to be used in excel file
func (rd *RelevantData) ReturnRelevantData(id string) *RelevantData {
	return rd.setFields(rd.retrieveRelevantData(id))
}

func (rd *RelevantData) setFields(pd, vc string) *RelevantData {
	rd.PD = pd
	rd.VC = vc
	return rd
}

func (rd *RelevantData) retrieveRelevantData(id string) (string, string) {
	client := youtube.Client{}

	video, err := client.GetVideo(id)
	Enil(err)
	return video.PublishDate.Format("2006-01-02"), video.ViewCount
}

// TODO
// Use Goroutines for multiple downloads
