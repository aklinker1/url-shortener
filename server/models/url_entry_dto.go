package models

import (
	"time"
)

type URLEntryDTO struct {
	ID        uint64     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	HashedID  string     `json:"hashedId"`
	URL       string     `json:"url"`
	Visits    uint64     `jsdon:"visits"`
}
