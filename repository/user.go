package repository

import (
	"github.com/Krisna20046/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(username string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUsersByRole(role string) ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, nil // Mengembalikan model.User kosong jika pengguna tidak ditemukan
		}
		return model.User{}, err // Mengembalikan error jika terjadi kesalahan lain
	}
	return user, nil
	// TODO: replace this
}
func (r *userRepository) GetUsersByRole(role string) ([]model.User, error) {
	var users []model.User
	err := r.db.Where("role = ?", role).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
