package model

type Playlist struct {
	PlaylistName  string  `json:"playlistName"`
	Protected     bool    `json:"protected"`
	Description   string  `json:"description"`
	Videos        []Video `json:"videos"`
	Id            string  `json:"_id"`
	CreatedAt     uint64  `json:"createdAt"`
	LastUpdatedAt uint64  `json:"lastUpdatedAt"`
}
