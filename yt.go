/*
Retrieve YouTube data
*/

package main

// https://thewebivore.com/using-replace-in-go-mod-to-point-to-your-local-module/

import (
	"fmt"
	"reflect"

	"github.com/kkdai/youtube/v2"
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
