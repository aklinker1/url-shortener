package models

type URLEntry struct {
	ID     string `gorm:"primaryKey"`
	URL    string `gorm:"unique"`
	Visits uint64
}
