package model

import "time"

type Profile struct {
	ID        uint       `json:"-";gorm:"primary_key"`
	DeletedAt *time.Time `json:"-";sql:"index"`

	Name          string         `json:"name"`
	BgColor       string         `json:"bgColor"`
	TextColor     string         `json:"textColor"`
	Subscriptions []Subscription `json:"subscriptions" gorm:"many2many:profile_subscription"`
}
