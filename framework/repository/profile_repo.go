package repositoryimpl

import (
	"github.com/Ceazzer/personal-website-server/core/domain/entity"
	"github.com/Ceazzer/personal-website-server/core/repository"
	"github.com/jinzhu/gorm"
)

func ProfileRepoFunc(db *gorm.DB) (*repository.ProfileRepo, error) {

	create := func(data *entity.Profile) (int64, error) {
		result := db.Create(data)
		return data.ID, result.Error
	}

	delete := func(id int64) (int64, error) {
		result := db.Delete(&entity.Profile{}, id)
		return id, result.Error
	}

	return &repository.ProfileRepo{
		Create: create,
		Delete: delete,
	}, nil
}
