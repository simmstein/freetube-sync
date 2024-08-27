package model

type Video struct {
	VideoId                  string `json:"videoId"`
	Title                    string `json:"title"`
	Author                   string `json:"author"`
	AuthorId                 string `json:"authorId"`
	Published                uint64 `json:"published"`
	Description              string `json:"description"`
	ViewCount                uint64 `json:"viewCount"`
	LengthSeconds            uint64 `json:"lengthSeconds"`
	WatchProgress            uint64 `json:"watchProgress"`
	TimeWatched              uint64 `json:"timeWatched"`
	LsLive                   bool   `json:"isLive"`
	Type                     string `json:"type"`
	Id                       string `json:"_id"`
	LastViewedPlaylistType   string `json:"lastViewedPlaylistType"`
	LastViewedPlaylistItemId string `json:"lastViewedPlaylistItemId"`
}
