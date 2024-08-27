package model

type Profile struct {
	Name          string         `json:"name"`
	BgColor       string         `json:"bgColor"`
	TextColor     string         `json:"textColor"`
	Subscriptions []Subscription `json:"subscriptions"`
}
