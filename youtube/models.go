package youtube

// RestResponse struct
type RestResponse struct {
	Items []Item `json:"items"`
}

// Item struct
type Item struct {
	ID ItemInfo `json:"id"`
}

// ItemInfo struct
type ItemInfo struct {
	VideoID string `json:"videoId"`
}
