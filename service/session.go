package service

import (
	"github.com/Krisna20046/model"
	repo "github.com/Krisna20046/repository"
)

type SessionService interface {
	GetSessionByUsername(username string) (model.Session, error)
}

type sessionService struct {
	sessionRepo repo.SessionRepository
}

func NewSessionService(sessionRepo repo.SessionRepository) *sessionService {
	return &sessionService{sessionRepo}
}

func (c *sessionService) GetSessionByUsername(username string) (model.Session, error) {
	session, err := c.sessionRepo.SessionAvailUsername(username)
	if err != nil {
		return model.Session{}, err
	}
	return session, nil
	 // TODO: replace this
}
