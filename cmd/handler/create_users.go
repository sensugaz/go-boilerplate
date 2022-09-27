package handler

import (
	"log"
	"net/http"

	"gitlab.bitkub.io/bo/gokub-boilerplate/cmd/domain/user"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUserHandler(c echo.Context) error {
	usr := &user.Request{}
	err := c.Bind(usr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(usr)
	if err != nil {
		return err
	}

	usrService, err := h.UserService.Create(&user.User{
		Name:  usr.Name,
		Email: usr.Email,
	})
	if err != nil {
		return err
	}

	// PLEASE REMOVE
	log.Println(usrService.Name)
	log.Println(usrService.Email)

	return c.JSON(http.StatusOK, usr)
}
