package handler

import (
	"github.com/Ceazzer/personal-website-server/core/domain/entity"
	"github.com/Ceazzer/personal-website-server/core/usecase"
	repositoryimpl "github.com/Ceazzer/personal-website-server/framework/repository"
	usecaseimpl "github.com/Ceazzer/personal-website-server/framework/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	CreateProfileHandler func(c echo.Context) error
	GetProfileHandler    func(c echo.Context) error
}

func NewProfileHandler(e *echo.Echo, db *gorm.DB) *ProfileHandler {
	profileRepo, err := repositoryimpl.ProfileRepoFunc(db)
	if err != nil {
		panic(err)
	}

	profileUsecase, err := usecaseimpl.ProfileUsecaseFunc(&usecase.ProfileUsecaseOpts{
		Repo: profileRepo,
	})
	if err != nil {
		panic(err)
	}

	cph := func(c echo.Context) error {
		profile := entity.Profile{}

		if err := c.Bind(profile); err != nil {
			return c.String(400, "Bad Request")
		}

		profileUsecase.CreateProfile(c.Request().Context(), &profile)
		return c.String(200, "Profile Create Handler")
	}

	gph := func(c echo.Context) error {
		return c.String(200, "Profile Get Handler")
	}

	return &ProfileHandler{
		CreateProfileHandler: cph,
		GetProfileHandler:    gph,
	}
}
