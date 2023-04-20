package main

import (
	"fmt"
	"github.com/AudDMusic/audd-go"
	"os"
)

func main() {



	client := audd.NewClient(os.Getenv("API_KEY"))
	file, err := os.Open("./music.mp3")

	if err != nil {
		panic(err)
	}
	result, err := client.Recognize(file, "apple_music,spotify", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s - %s.\nTimecode: %s, album: %s. ℗ %s, %s\n\n"+
		"Listen: %s\nOr directly on:\n- Apple Music: %s, \n- Spotify: %s",
		result.Artist, result.Title, result.Timecode, result.Album,
		result.Label, result.ReleaseDate, result.SongLink,
		result.AppleMusic.URL, result.Spotify.ExternalUrls.Spotify)
}