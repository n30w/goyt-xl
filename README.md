# Go YouTube-Excel
YouTube metadata utility that outputs things into excel format, written in Go. The original metadata utility was written in Python. I'm rewriting this one in Go because of the flexibility and native binary support, instead of having to create a virtual environment everytime the program needs to be run. Also, Go is just awesome.

## Packages
As of right now, this project uses these packages:
- github.com/kkdai/youtube/v2
- github.com/xuri/excelize/v2

I'm using my own edited version of the youtube package. This can be seen via the initial commit:

Please note: I edited the video.go files to include a new field for type
Video struct. This adds the ViewCount field, taken from
response_data.go. After adding this field, I refractored the method
extractDataFromPlayerResponse and added:

v.ViewCount = prData.VideoDetails.ViewCount

This lets me pull view count

## Useful Links
### Github
- [How to add Remote Repo to Github](https://articles.assembla.com/en/articles/1136998-how-to-add-a-new-remote-to-your-git-repo)
- [Your First Time With Git and Github](https://kbroman.org/github_tutorial/pages/first_time.html)
- [Start a new Git Repository](https://kbroman.org/github_tutorial/pages/init.html)

### Go Related
- [Using Replace in Go.Mod to point to local module](https://thewebivore.com/using-replace-in-go-mod-to-point-to-your-local-module/)
- [Importing local Go Packages](https://linguinecode.com/post/how-to-import-local-files-packages-in-golang)
- [Using Go Packages](https://levelup.gitconnected.com/using-modules-and-packages-in-go-36a418960556)
- [How to append to slice](https://go.dev/tour/moretypes/15)
- [Modifying a slice]( https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang)