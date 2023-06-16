package repository

import (
	"github.com/Krisna20046/model"

	"gorm.io/gorm"
)

type MakamRepository interface {
	Store(task *model.DataMakam) error
	Update(id int, task *model.DataMakam) error
	Delete(id int) error
	GetByID(id int) (*model.DataMakam, error)
	GetList() ([]model.DataMakam, error)
}

type makamRepository struct {
	db *gorm.DB
}

func NewMakamRepo(db *gorm.DB) *makamRepository {
	return &makamRepository{db}
}

func (t *makamRepository) Store(makam *model.DataMakam) error {
	err := t.db.Create(makam).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *makamRepository) Update(id int, makam *model.DataMakam) error {
	err := t.db.Model(&model.DataMakam{}).Where("id = ?", makam.ID).Updates(makam).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t *makamRepository) Delete(id int) error {
	err := t.db.Delete(&model.DataMakam{}, id).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t *makamRepository) GetByID(id int) (*model.DataMakam, error) {
	var makam model.DataMakam
	err := t.db.First(&makam, id).Error
	if err != nil {
		return nil, err
	}

	return &makam, nil
}

func (t *makamRepository) GetList() ([]model.DataMakam, error) {
	var makams []model.DataMakam
	err := t.db.Find(&makams).Error
	if err != nil {
		return nil, err
	}
	return makams, nil // TODO: replace this
}