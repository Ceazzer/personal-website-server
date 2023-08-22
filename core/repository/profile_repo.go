package repository

import (
	"github.com/Ceazzer/personal-website-server/core/domain"
	"github.com/jinzhu/gorm"
)

// Functions
type profileDeleteFunc func(id int64) (int64, error)
type profileCreateFunc func(data *domain.Profile) (int64, error)

type ProfileRepoFunc func(db *gorm.DB) (*RepoType, error)

type RepoType struct {
	Create profileCreateFunc
	Delete profileDeleteFunc
}

func New(db *gorm.DB) (*RepoType, error) {

	create := func(data *domain.Profile) (int64, error) {
		result := db.Create(data)
		return data.ID, result.Error
	}

	delete := func(id int64) (int64, error) {
		result := db.Delete(&domain.Profile{}, id)
		return id, result.Error
	}

	return &RepoType{
		Create: create,
		Delete: delete,
	}, nil
}
