package model

type Profile struct {
	ID uint `json:"-" gorm:"primary_key"`

	Name          string         `json:"name"`
	BgColor       string         `json:"bgColor"`
	TextColor     string         `json:"textColor"`
	Subscriptions []Subscription `json:"subscriptions" gorm:"many2many:profile_subscription"`
	RemoteId      string         `json:"_id"`
}
