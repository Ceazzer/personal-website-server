package repositoryimpl

import (
	"github.com/Ceazzer/personal-website-server/core/domain/entity"
	"github.com/Ceazzer/personal-website-server/core/repository"
)

// MessageRepoImpl returns a MessageRepoResponse struct
// with the FindByID and Create functions implemented
// with the logic from the MessageRepoFunc function
func MessageRepoImpl() (*repository.MessageRepoResponse, error) {

	// Create Message Function
	create := func(data *entity.Message) (string, error) {
		return "", nil
	}

	// Find Message By ID Function
	findByID := func(id int) (*entity.Message, error) {
		return nil, nil
	}

	// Return Repo
	return &repository.MessageRepoResponse{
		Create:   create,
		FindByID: findByID,
	}, nil
}
