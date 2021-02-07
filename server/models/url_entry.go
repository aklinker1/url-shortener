package models

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type URLEntry struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	URL       string         `gorm:"index"`
	Visits    uint64
}

func (model *URLEntry) ToDTO() *URLEntryDTO {
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	return &URLEntryDTO{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: deletedAt,
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
