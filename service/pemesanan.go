package service

import (
	"github.com/Krisna20046/model"
	repo "github.com/Krisna20046/repository"
)

type PemesananService interface {
	Store(pesan *model.Pemesanan) error
	Update(id int, Pemesanan *model.Pemesanan) error
	Delete(id int) error
	GetByID(id int) (*model.Pemesanan, error)
	GetList() ([]model.Pemesanan, error)
}

type pemesananService struct {
	pemesananRepository repo.PemesananRepository
}

func NewPemesananService(pemesananRepository repo.PemesananRepository) PemesananService {
	return &pemesananService{pemesananRepository}
}

func (c *pemesananService) Store(pesan *model.Pemesanan) error {
	err := c.pemesananRepository.Store(pesan)
	if err != nil {
		return err
	}

	return nil
}

func (s *pemesananService) Update(id int, pesan *model.Pemesanan) error {
	err := s.pemesananRepository.Update(id, pesan)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *pemesananService) Delete(id int) error {
	err := s.pemesananRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *pemesananService) GetByID(id int) (*model.Pemesanan, error) {
	pesan, err := s.pemesananRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return pesan, nil
}

func (s *pemesananService) GetList() ([]model.Pemesanan, error) {
	pemesanans, err := s.pemesananRepository.GetList()
	if err != nil {
		return nil, err
	}
	return pemesanans, nil // TODO: replace this
}
