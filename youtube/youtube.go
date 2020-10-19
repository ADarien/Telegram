package youtube

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const YOUTUBE_SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
const YOUTUBE_API_TOKEN = "AIzaSyDNbiNEt30YxlbQb8ZHkP_3j48WH69GboY"
const YOUTUBE_VIDEO_URL = "https://www.youtube.com/watch?v="

// GET https://www.googleapis.com/youtube/v3/search?part=id&channelId=UCY6zVRa3Km52bsBmpyQnk6A&maxResults=1&order=date&key=[YOUR_API_KEY] HTTP/1.1

// Authorization: Bearer [YOUR_ACCESS_TOKEN]
// Accept: application/json

// GetLastVideo func
func GetLastVideo(channelURL string) (string, error) {
	items, err := retrieveVideos(channelURL)
	if err != nil {
		return "", err
	}
	if len(items) < 1 {
		return "", errors.New("No video found")
	}
	return YOUTUBE_VIDEO_URL + items[0].ID.VideoID, nil
}

func retrieveVideos(channeURL string) ([]Item, error) {
	req, err := makeRequest(channeURL, 1)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Items, nil
}

func makeRequest(channeURL string, maxResults int) (*http.Request, error) {
	lastSlashIndex := strings.LastIndex(channeURL, "/")
	channelID := channeURL[lastSlashIndex+1:]
	req, err := http.NewRequest("GET", YOUTUBE_SEARCH_URL, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("part", "id")
	query.Add("channelId", channelID)
	query.Add("maxResults", strconv.Itoa(maxResults))
	query.Add("order", "date")
	query.Add("key", YOUTUBE_API_TOKEN)
	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())
	return req, nil
}
