package repositoryimpl

import (
	"github.com/Ceazzer/personal-website-server/core/domain/dto"
	"github.com/Ceazzer/personal-website-server/core/repository"
)

// MessageRepoImpl returns a MessageRepoResponse struct
// with the FindByID and Create functions implemented
// with the logic from the MessageRepoFunc function
func MessageRepoImpl() (*repository.MessageRepoResponse, error) {
	repo := &repository.MessageRepoResponse{}

	// Create Message Function
	repo.Create = func(data *dto.MessageDTO) (string, error) {
		return "", nil
	}

	// Find Message By ID Function
	repo.FindByID = func(id int) (*dto.MessageDTO, error) {
		return nil, nil
	}

	// Return Repo
	return repo, nil
}
