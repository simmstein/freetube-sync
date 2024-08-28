package model

import (
	"time"
)

type Pull struct {
	ID uint `json:"-";gorm:"primary_key"`

	Hostname string
	Database string
	PullAt   time.Time
}
