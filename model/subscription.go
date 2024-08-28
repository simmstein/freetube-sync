package model

import "time"

type Subscription struct {
	ID        uint       `json:"-";gorm:"primary_key"`
	DeletedAt *time.Time `json:"-";sql:"index"`

	Profiles  []Profile `gorm:"many2many:profile_subscription"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Thumbnail string    `json:"thumbnail"`
}
