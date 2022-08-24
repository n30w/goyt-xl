/*
Retrieve YouTube data
*/

package main

import (
	"fmt"
	"reflect"

	"github.com/kkdai/youtube/v2"
	// "github.com/n30w/goyt-xl/youtube"
)

func ExampleClient() {
	videoID := "BaW_jenozKc"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)

	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.TypeOf(video.PublishDate))
	fmt.Println(video.ViewCount)
}

/*
TODO:
Get video view count via playerResponseData struct
*/
