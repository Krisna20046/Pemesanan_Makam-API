package service

import (
	"github.com/Krisna20046/model"
	repo "github.com/Krisna20046/repository"
)

type JenazahService interface {
	Store(jenazah *model.DataJenazah) error
	Update(id int, DataJenazah *model.DataJenazah) error
	Delete(id int) error
	GetByID(id int) (*model.DataJenazah, error)
	GetList() ([]model.DataJenazah, error)
}

type jenazahService struct {
	jenazahRepository repo.JenazahRepository
}

func NewJenazahService(jenazahRepository repo.JenazahRepository) JenazahService {
	return &jenazahService{jenazahRepository}
}

func (c *jenazahService) Store(jenazah *model.DataJenazah) error {
	err := c.jenazahRepository.Store(jenazah)
	if err != nil {
		return err
	}

	return nil
}

func (s *jenazahService) Update(id int, jenazah *model.DataJenazah) error {
	err := s.jenazahRepository.Update(id, jenazah)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *jenazahService) Delete(id int) error {
	err := s.jenazahRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *jenazahService) GetByID(id int) (*model.DataJenazah, error) {
	jenazah, err := s.jenazahRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return jenazah, nil
}

func (s *jenazahService) GetList() ([]model.DataJenazah, error) {
	jenazahs, err := s.jenazahRepository.GetList()
	if err != nil {
		return nil, err
	}
	return jenazahs, nil // TODO: replace this
}
