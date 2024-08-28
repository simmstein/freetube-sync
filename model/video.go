package model

import (
	"time"

	"gorm.io/gorm"
)

type WatchedVideo struct {
	ID        uint           `json:"-" gorm:"primary_key"`
	DeletedAt gorm.DeletedAt `json:"-" sql:"index"`

	VideoId                  string  `json:"videoId"`
	Title                    string  `json:"title"`
	Author                   string  `json:"author"`
	AuthorId                 string  `json:"authorId"`
	Published                uint64  `json:"published"`
	Description              string  `json:"description"`
	ViewCount                uint64  `json:"viewCount"`
	LengthSeconds            uint64  `json:"lengthSeconds"`
	WatchProgress            uint64  `json:"watchProgress"`
	TimeWatched              uint64  `json:"timeWatched"`
	IsLive                   bool    `json:"isLive"`
	Type                     string  `json:"type"`
	RemoteId                 string  `json:"_id"`
	LastViewedPlaylistType   string  `json:"lastViewedPlaylistType"`
	LastViewedPlaylistItemId *string `json:"lastViewedPlaylistItemId"`
}

type PlaylistVideo struct {
	ID        uint       `gorm:"primary_key"`
	DeletedAt *time.Time `json:"-" sql:"index"`

	PlaylistID     uint
	VideoId        string `json:"videoId"`
	Title          string `json:"title"`
	AuthorId       string `json:"authorId"`
	LengthSeconds  uint64 `json:"lengthSeconds"`
	TimeWatched    uint64 `json:"timeWatched"`
	TimeAdded      uint64 `json:"timeAdded"`
	PlaylistItemId string `json:"playlistItemId"`
	Type           string `json:"type"`
}
