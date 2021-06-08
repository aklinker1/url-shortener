package repos

import (
	"gorm.io/gorm"

	"github.com/aklinker1/url-shortener/server/models"
)

type urlEntryRepo struct {
	db *gorm.DB
}

var URLEntryRepo = &urlEntryRepo{}

func (repo *urlEntryRepo) Init(db *gorm.DB) {
	URLEntryRepo.db = db
}

func (repo *urlEntryRepo) Create(url string) (*models.URLEntry, error) {
	entry := &models.URLEntry{
		URL: url,
	}
	err := repo.db.Create(entry).Error
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (repo *urlEntryRepo) Delete(entry *models.URLEntry) error {
	return repo.db.Delete(entry).Error
}

func (repo *urlEntryRepo) Read(id int64) (*models.URLEntry, error) {
	model := &models.URLEntry{}
	err := repo.db.First(model, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *urlEntryRepo) Update(entry *models.URLEntry, url string) (*models.URLEntry, error) {
	entry.URL = url
	err := repo.db.Save(entry).Error
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (repo *urlEntryRepo) UpdateVisits(entry *models.URLEntry) (*models.URLEntry, error) {
	entry.Visits++
	err := repo.db.Save(entry).Error
	if err != nil {
		return nil, err
	}
	return entry, nil
}

// func (repo *urlEntryRepo) Search(url string) (*models.URLEntry, error) {
// 	panic("NOT IMPLEMENTED")
// }

func (repo *urlEntryRepo) List(pagination *models.Pagination, search string) ([]*models.URLEntry, error) {
	models := []*models.URLEntry{}
	query := repo.db.
		Limit(pagination.Limit()).
		Offset(pagination.Offset()).
		Order("created_at DESC")
	var err error
	if search != "" {
		err = query.Find(&models, "url LIKE ?", "%"+search+"%").Error
	} else {
		err = query.Find(&models).Error
	}
	if err != nil {
		return nil, err
	}
	return models, nil
}
