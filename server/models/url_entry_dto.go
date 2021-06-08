package models

import (
	"time"
)

type URLEntryDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	URL       string    `json:"url"`
	Visits    uint64    `jsdon:"visits"`
}
