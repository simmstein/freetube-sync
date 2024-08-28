package model

import "time"

type Playlist struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`

	Hostname        string          `json:"-"`
	PlaylistName    string          `json:"playlistName"`
	Protected       bool            `json:"protected"`
	Description     string          `json:"description"`
	Videos          []PlaylistVideo `json:"videos"`
	RemoteId        string          `json:"_id"`
	RemoteCreatedAt uint64          `json:"createdAt"`
	LastUpdatedAt   uint64          `json:"lastUpdatedAt"`
}
