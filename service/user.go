package service

import (
	"github.com/Krisna20046/model"
	repo "github.com/Krisna20046/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
	GetUsersByRole(role string) ([]model.User, error)
}

type userService struct {
	userRepo     repo.UserRepository
	sessionsRepo repo.SessionRepository
}

func NewUserService(userRepository repo.UserRepository, sessionsRepo repo.SessionRepository) UserService {
	return &userService{userRepository, sessionsRepo}
}

func (s *userService) GetUsersByRole(role string) ([]model.User, error) {
	users, err := s.userRepo.GetUsersByRole(role)
	if err != nil {
		return nil, err
	}
	return users, nil
}


func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		return *user, err
	}

	if dbUser.Username != "" || dbUser.ID != 0 {
		return *user, errors.New("username already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Login(user *model.User) (token *string, err error) {
	dbUser, err := s.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	if dbUser.Username == "" || dbUser.ID == 0 {
		return nil, errors.New("user not found")
	}

	if user.Password != dbUser.Password {
		return nil, errors.New("wrong username or password")
	}

	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &model.Claims{
		Username: user.Username,
		Role:     dbUser.Role, // Set nilai peran (role) dari user yang ditemukan di database
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(model.JwtKey)
	if err != nil {
		return nil, err
	}

	session := model.Session{
		Token:    tokenString,
		Username: user.Username,
		Expiry:   expirationTime,
	}

	_, err = s.sessionsRepo.SessionAvailUsername(session.Username)
	if err != nil {
		err = s.sessionsRepo.AddSessions(session)
	} else {
		err = s.sessionsRepo.UpdateSessions(session)
	}

	return &tokenString, nil
}



