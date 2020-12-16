package api

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TrackInfo represents track information
type TrackInfo struct {
	Title       string
	Artist      string
	Genres      []string
	DownloadUrl string
}

/*func GetTracksInfoByGenre(genre string, pageNum int) ([]*TrackInfo, error) {
	url := constructUrl(
		"https",
		hostName,
		"music/"+genre,
		map[string]string{
			"download": "1",
			"page":     strconv.Itoa(pageNum)})

	doc, err := goquery.NewDocument(url.String())
	if err != nil {
		return nil, err
	}

	var tracks []*TrackInfo

	doc.Find("div.title > a").Each(func(i int, s *goquery.Selection) {
		trackPageURL := s.AttrOr("href", "")
		if trackPageURL == "" {
			return
		}

		trackInfo, err := parsePage(trackPageURL)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(*trackInfo)

		tracks = append(tracks, trackInfo)
	})

	return tracks, nil
}*/

func parsePage(url string) (*TrackInfo, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	title := doc.Find("span.file_title ").First().Text()

	artist := doc.Find("div.dj_head_text > a").First().Text()

	var genres []string
	doc.Find("span.styles > a").Each(func(i int, s *goquery.Selection) {
		genres = append(genres, s.Text())
	})

	downloadUrl := doc.Find("a#download_flasher").First().AttrOr("href", "")

	trackInfo := &TrackInfo{
		Title:       title,
		Artist:      artist,
		Genres:      genres,
		DownloadUrl: downloadUrl}

	return trackInfo, nil
}

// RandomTrackInfoByGenre returns information about random track with specified genre
func RandomTrackInfoByGenre(genre string) (*TrackInfo, error) {
	const maxTrackNum = 1000

	trackNum := rand.Intn(maxTrackNum)

	pageNum := trackNum/20 + 1
	tNum := trackNum - (pageNum-1)*20 - 1

	url := constructUrl(
		"https",
		hostName,
		"music/"+genre,
		map[string]string{
			"download": "1",
			"page":     strconv.Itoa(pageNum)})

	doc, err := goquery.NewDocument(url.String())
	if err != nil {
		return nil, err
	}

	trackPageUrl := doc.Find("div.title > a").Eq(tNum).AttrOr("href", "")
	if trackPageUrl == "" {
		return nil, errors.New("trackPageUrl is empty")
	}

	return parsePage(trackPageUrl)
}

// Search returns tracks information by specified search query
func Search(query string) (result []*TrackInfo, err error) {
	type File struct {
		Kind string
		Html string
	}
	type Results []File
	type Response struct {
		Results Results
	}

	url := constructUrl(
		"https",
		hostName,
		"search",
		map[string]string{
			"mode":      "file",
			"searchfor": query,
			"sortby":    "relevance",
			"period":    "all",
			"page":      "1",
			"results":   "1"})

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	for _, item := range response.Results {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(item.Html))
		if err != nil {
			return nil, err
		}
		trackInfo, err := parsePage(doc.Find("div.title > a").First().AttrOr("href", ""))
		if err != nil {
			return nil, err
		}
		result = append(result, trackInfo)
	}

	return result, nil
}
