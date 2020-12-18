package api

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Genre represents genre information
type Genre struct {
	Name string
	Code string
}

// Genres holds cached list of available genres
var Genres []Genre

func UpdateGenres() error {
	var err error

	Genres, err = updateGenreList()
	if err != nil {
		return fmt.Errorf("get genres list failed: %w", err)
	}

	return nil
}

func updateGenreList() ([]Genre, error) {
	url := "https://promodj.com/music"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	var genres []Genre
	doc.Find("div.styles_tagcloud > a").Each(func(i int, s *goquery.Selection) {
		genres = append(genres, Genre{s.Text(), strings.TrimPrefix(s.AttrOr("href", ""), "/music/")})
	})

	return genres, nil
}

// GetM3uPlaylist returns playlist of random track genres. Hostname is specified
// with urlPrefix value.
func GetM3uPlaylist(urlPrefix string) []byte {
	b := new(bytes.Buffer)

	fmt.Fprint(b, "#EXTM3U\n")

	for _, genre := range Genres {
		fmt.Fprintf(b, "#EXTINF:-1, %s\n", genre.Name)
		fmt.Fprintf(b, "%s%s\n", urlPrefix, genre.Code)
	}

	return b.Bytes()
}
