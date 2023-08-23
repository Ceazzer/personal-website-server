package usecase

type ProfileUsecase interface{}

type profileUsecase struct{}

func NewProfileUsecase() ProfileUsecase {
	return &profileUsecase{}
}

func (u *profileUsecase) CreateProfile() {}

func (u *profileUsecase) DeleteProfile() {}
