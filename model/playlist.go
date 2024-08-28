package model

import "time"

type Playlist struct {
	ID        uint       `json:"-";gorm:"primary_key"`
	DeletedAt *time.Time `json:"-";sql:"index"`

	PlaylistName  string          `json:"playlistName"`
	Protected     bool            `json:"protected"`
	Description   string          `json:"description"`
	Videos        []PlaylistVideo `json:"videos"`
	Id            string          `json:"_id"`
	CreatedAt     uint64          `json:"createdAt"`
	LastUpdatedAt uint64          `json:"lastUpdatedAt"`
}
