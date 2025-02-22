package handler

import (
	"net/http"
	"repocon/usecase"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	GetUser(c echo.Context) error
}

type userHandler struct {
	uu usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) IUserHandler {
	return &userHandler{uu}
}

func (uh *userHandler) GetUser(c echo.Context) error {

	claims, ok := c.Get("user").(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}
	id := c.Param("id")

	tokenID, ok := claims["sub"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	if id != tokenID {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	user, err := uh.uu.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user")
	}
	return c.JSON(http.StatusOK, user)
}
