package repos

import (
	"errors"

	"gorm.io/gorm"

	"github.com/aklinker1/url-shortener/server/models"
	"github.com/aklinker1/url-shortener/server/utils"
)

type urlEntryRepo struct {
	db *gorm.DB
}

var URLEntryRepo = &urlEntryRepo{}
var ID_CHARACTER_SET = "abcdefghijkmnpqrstuvwxyz" + "ABCDEFGHJKLMNPQRSTUVWXYZ" + "0123456789"

func (repo *urlEntryRepo) Init(db *gorm.DB) {
	URLEntryRepo.db = db
}

func (repo *urlEntryRepo) nextID() (string, error) {
	for idLength := 1; idLength < 8; idLength++ {
		for attempts := 0; attempts < 4; attempts++ {
			newID := utils.RandomStringWithCharset(idLength, ID_CHARACTER_SET)
			_, err := repo.Read(newID)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return newID, nil
			}
		}
	}

	return "", errors.New("Could not generate a new id")
}

func (repo *urlEntryRepo) Create(url string) (*models.URLEntry, error) {
	id, err := repo.nextID()
	if err != nil {
		return nil, err
	}
	entry := &models.URLEntry{
		ID:  id,
		URL: url,
	}

	return entry, repo.db.Create(entry).Error
}

func (repo *urlEntryRepo) Delete(entry *models.URLEntry) error {
	return repo.db.Delete(entry).Error
}

func (repo *urlEntryRepo) Read(id string) (*models.URLEntry, error) {
	model := &models.URLEntry{}
	err := repo.db.First(model, "id = ?", id).Error
	return model, err
}

func (repo *urlEntryRepo) Update(entry *models.URLEntry, url string) (*models.URLEntry, error) {
	entry.URL = url
	err := repo.db.Save(entry).Error
	return entry, err
}

func (repo *urlEntryRepo) UpdateVisits(entry *models.URLEntry) (*models.URLEntry, error) {
	entry.Visits++
	err := repo.db.Save(entry).Error
	return entry, err
}

// func (repo *urlEntryRepo) Search(url string) (*models.URLEntry, error) {
// 	panic("NOT IMPLEMENTED")
// }

func (repo *urlEntryRepo) List(page, size int) ([]*models.URLEntry, error) {
	models := []*models.URLEntry{}
	err := repo.db.Find(&models).Limit(size).Offset(page * size).Error
	return models, err
}
