package models

import (
	"time"
)

type URLEntry struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	URL       string `gorm:"unique"`
	Visits    uint64
}
