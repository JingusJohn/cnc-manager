package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/JingusJohn/kyle-cnc/backend/models"
)

type AuthController struct{}

func CreateAuthController() *AuthController {
	return &AuthController{}
}

func (controller *AuthController) Register(app *echo.Echo) {
	group := app.Group("/auth")
	group.POST("/login", login)
}

func login(c echo.Context) error {
	// parse the body of the request (Status 400 if failure)
	var credentials models.UserLogin
	err := c.Bind(&credentials)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid credentials")
	}

	return c.String(http.StatusOK, "working")
}
