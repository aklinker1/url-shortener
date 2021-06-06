package models

import (
	"strconv"
	"time"
)

type URLEntry struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	URL       string    `gorm:"unique" json:"url"`
	Visits    uint64    `json:"visits"`
}

func (model *URLEntry) ToDTO() *URLEntryDTO {
	return &URLEntryDTO{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		URL:       model.URL,
		Visits:    model.Visits,
		HashedID:  strconv.FormatInt(int64(model.ID), 32),
	}
}

func ToDTOs(entries []*URLEntry) []*URLEntryDTO {
	dtos := []*URLEntryDTO{}
	for _, entry := range entries {
		dtos = append(dtos, entry.ToDTO())
	}
	return dtos
}
