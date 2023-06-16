package service

import (
	"github.com/Krisna20046/model"
	repo "github.com/Krisna20046/repository"
)

type MakamService interface {
	Store(makam *model.DataMakam) error
	Update(id int, DataMakam *model.DataMakam) error
	Delete(id int) error
	GetByID(id int) (*model.DataMakam, error)
	GetList() ([]model.DataMakam, error)
}

type makamService struct {
	makamRepository repo.MakamRepository
}

func NewMakamService(makamRepository repo.MakamRepository) MakamService {
	return &makamService{makamRepository}
}

func (c *makamService) Store(makam *model.DataMakam) error {
	err := c.makamRepository.Store(makam)
	if err != nil {
		return err
	}

	return nil
}

func (s *makamService) Update(id int, makam *model.DataMakam) error {
	err := s.makamRepository.Update(id, makam)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *makamService) Delete(id int) error {
	err := s.makamRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *makamService) GetByID(id int) (*model.DataMakam, error) {
	makam, err := s.makamRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return makam, nil
}

func (s *makamService) GetList() ([]model.DataMakam, error) {
	makams, err := s.makamRepository.GetList()
	if err != nil {
		return nil, err
	}
	return makams, nil // TODO: replace this
}
