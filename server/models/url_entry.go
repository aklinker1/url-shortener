package models

import (
	"strconv"
	"time"
)

type URLEntry struct {
	ID        uint64 `gorm:"primaryKey" `
	CreatedAt time.Time
	URL       string `gorm:"unique" `
	Visits    uint64
}

func (model *URLEntry) ToDTO() *URLEntryDTO {
	return &URLEntryDTO{
		ID:        strconv.FormatInt(int64(model.ID), 32),
		CreatedAt: model.CreatedAt,
		URL:       model.URL,
		Visits:    model.Visits,
	}
}

func ToDTOs(entries []*URLEntry) []*URLEntryDTO {
	dtos := []*URLEntryDTO{}
	for _, entry := range entries {
		dtos = append(dtos, entry.ToDTO())
	}
	return dtos
}
