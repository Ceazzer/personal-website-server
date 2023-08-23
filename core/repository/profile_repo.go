package repository

import (
	"github.com/Ceazzer/personal-website-server/core/domain"
	"github.com/jinzhu/gorm"
)

type ProfileRepository interface {
	Create(data *domain.Profile) (int64, error)
	GetAll() ([]*domain.Profile, error)
	GetByID(id int64) (*domain.Profile, error)
	Delete(id int64) (int64, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository() ProfileRepository {
	return &profileRepository{}
}

func (r *profileRepository) Create(data *domain.Profile) (int64, error) {
	result := r.db.Create(data)
	return data.ID, result.Error
}

func (r *profileRepository) GetAll() ([]*domain.Profile, error) {
	var profiles []*domain.Profile
	result := r.db.Find(&profiles)
	return profiles, result.Error
}

func (r *profileRepository) GetByID(id int64) (*domain.Profile, error) {
	var profile domain.Profile
	result := r.db.First(&profile, id)
	return &profile, result.Error
}

func (r *profileRepository) Delete(id int64) (int64, error) {
	result := r.db.Delete(&domain.Profile{}, id)
	return id, result.Error
}
