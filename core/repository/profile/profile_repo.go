package profilerepository

import (
	"github.com/Ceazzer/personal-website-server/core/domain/entity"

	"github.com/jinzhu/gorm"
)

// Functions
type profileCreateFunc func(data *entity.Profile) (int64, error)
type profileDeleteFunc func(id int64) (int64, error)

type RepoType struct {
	Create profileCreateFunc
	Delete profileDeleteFunc
}

type ProfileRepoFunc func() (*RepoType, error)

func New(db *gorm.DB) (*RepoType, error) {

	create := func(data *entity.Profile) (int64, error) {
		result := db.Create(data)
		return data.ID, result.Error
	}

	delete := func(id int64) (int64, error) {
		result := db.Delete(&entity.Profile{}, id)
		return id, result.Error
	}

	return &RepoType{
		Create: create,
		Delete: delete,
	}, nil
}
