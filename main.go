package main

import (
	"fmt"
	"github.com/AudDMusic/audd-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {


	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/upload", func(c *gin.Context) {

		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		err := c.SaveUploadedFile(file, "./musics/mm.mp3")

		if err != nil {
			panic(err)
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})

	router.Run(":8080")

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