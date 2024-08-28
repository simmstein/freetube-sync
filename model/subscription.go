package model

type Subscription struct {
	ID uint `json:"-" gorm:"primary_key"`

	Profiles  []Profile `gorm:"many2many:profile_subscription"`
	RemoteId  string    `json:"id"`
	Name      string    `json:"name"`
	Thumbnail string    `json:"thumbnail"`
}
