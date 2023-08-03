package handler

import "github.com/labstack/echo/v4"

type ProfileHandler struct {
	CreateProfileHandler func(c echo.Context) error
	GetProfileHandler    func(c echo.Context) error
}

func NewProfileHandler(e *echo.Echo) *ProfileHandler {

	cph := func(c echo.Context) error {
		return nil
	}

	gph := func(c echo.Context) error {
		return nil
	}

	return &ProfileHandler{
		CreateProfileHandler: cph,
		GetProfileHandler:    gph,
	}
}
