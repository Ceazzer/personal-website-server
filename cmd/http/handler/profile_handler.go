package handler

import "github.com/labstack/echo/v4"

type ProfileHandlerResponse struct {
	CreateProfileHandler func(c echo.Context) error
	GetProfileHandler    func(c echo.Context) error
}

func ProfileHandler(e *echo.Echo) *ProfileHandlerResponse {

	cph := func(c echo.Context) error {

		return c.String(200, "Profile Create Handler")
	}

	gph := func(c echo.Context) error {
		return c.String(200, "Profile Get Handler")
	}

	return &ProfileHandlerResponse{
		CreateProfileHandler: cph,
		GetProfileHandler:    gph,
	}
}
