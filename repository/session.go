package repository

import (
	"errors"
	"time"

	"github.com/Krisna20046/model"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailUsername(username string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	GetSessionsByRole(role string) ([]model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *sessionsRepo {
	return &sessionsRepo{db}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	err := u.db.Create(&session).Error
	return err // TODO: replace this
}

func (u *sessionsRepo) DeleteSession(token string) error {
	err := u.db.Where("token = ?", token).Delete(&model.Session{}).Error
	return err // TODO: replace this
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	err := u.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session).Error
	return err
	// TODO: replace this
}

func (u *sessionsRepo) SessionAvailUsername(username string) (model.Session, error) {
	var session model.Session
	err := u.db.Where("username = ?", username).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Session{}, errors.New("session not found")
		}
		return model.Session{}, err
	}

	return session, nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := u.db.Where("token = ?", token).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Session{}, errors.New("session not found")
		}
		return model.Session{}, err
	}

	return session, nil // TODO: replace this
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}

func (s *sessionsRepo) GetSessionsByRole(role string) ([]model.Session, error) {
	var sessions []model.Session
	err := s.db.Where("role = ?", role).Find(&sessions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []model.Session{}, errors.New("sessions not found")
		}
		return []model.Session{}, err
	}

	return sessions, nil
}
