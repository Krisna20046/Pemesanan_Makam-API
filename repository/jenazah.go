package repository

import (
	"github.com/Krisna20046/model"

	"gorm.io/gorm"
)

type JenazahRepository interface {
	Store(task *model.DataJenazah) error
	Update(id int, task *model.DataJenazah) error
	Delete(id int) error
	GetByID(id int) (*model.DataJenazah, error)
	GetList() ([]model.DataJenazah, error)
}

type jenazahRepository struct {
	db *gorm.DB
}

func NewJenazahRepo(db *gorm.DB) *jenazahRepository {
	return &jenazahRepository{db}
}

func (t *jenazahRepository) Store(jenazah *model.DataJenazah) error {
	err := t.db.Create(jenazah).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *jenazahRepository) Update(id int, jenazah *model.DataJenazah) error {
	err := t.db.Model(&model.DataJenazah{}).Where("id = ?", jenazah.ID).Updates(jenazah).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t *jenazahRepository) Delete(id int) error {
	err := t.db.Delete(&model.DataJenazah{}, id).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t *jenazahRepository) GetByID(id int) (*model.DataJenazah, error) {
	var jenazah model.DataJenazah
	err := t.db.First(&jenazah, id).Error
	if err != nil {
		return nil, err
	}

	return &jenazah, nil
}

func (t *jenazahRepository) GetList() ([]model.DataJenazah, error) {
	var jenazahs []model.DataJenazah
	err := t.db.Find(&jenazahs).Error
	if err != nil {
		return nil, err
	}
	return jenazahs, nil // TODO: replace this
}