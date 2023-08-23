package repository

import (
	"github.com/Ceazzer/personal-website-server/core/domain"
	"github.com/jinzhu/gorm"
)

type ProfileRepository interface {
	Create(data *domain.Profile) (int64, error)
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

func (r *profileRepository) Delete(id int64) (int64, error) {
	result := r.db.Delete(&domain.Profile{}, id)
	return id, result.Error
}
