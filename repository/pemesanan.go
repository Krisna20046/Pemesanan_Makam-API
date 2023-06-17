package repository

import (
	"github.com/Krisna20046/model"

	"gorm.io/gorm"
)

type PemesananRepository interface {
	Store(task *model.Pemesanan) error
	Update(id int, task *model.Pemesanan) error
	Delete(id int) error
	GetByID(id int) (*model.Pemesanan, error)
	GetList() ([]model.Pemesanan, error)
}

type pemesananRepository struct {
	db *gorm.DB
}

func NewPemesananRepo(db *gorm.DB) *pemesananRepository {
	return &pemesananRepository{db}
}

func (t *pemesananRepository) Store(pesan *model.Pemesanan) error {
	err := t.db.Create(pesan).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *pemesananRepository) Update(id int, pesan *model.Pemesanan) error {
	err := t.db.Model(&model.Pemesanan{}).Where("id = ?", pesan.ID).Updates(pesan).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t *pemesananRepository) Delete(id int) error {
	err := t.db.Delete(&model.Pemesanan{}, id).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t *pemesananRepository) GetByID(id int) (*model.Pemesanan, error) {
	var pesan model.Pemesanan
	err := t.db.First(&pesan, id).Error
	if err != nil {
		return nil, err
	}

	return &pesan, nil
}

func (t *pemesananRepository) GetList() ([]model.Pemesanan, error) {
	var pemesanans []model.Pemesanan
	err := t.db.Find(&pemesanans).Error
	if err != nil {
		return nil, err
	}
	return pemesanans, nil // TODO: replace this
}
