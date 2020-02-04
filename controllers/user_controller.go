package controllers

import (
	"errors"
	"go_api_boilerplate/services/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetById(*gin.Context)
}

type userController struct {
	us user_service.UserService
}

func NewUserController(us user_service.UserService) UserController {
	return &userController{
		us: us,
	}
}

func getUserId(userIdParam string) (int, error) {
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		return 0, errors.New("user id should be a number")
	}
	return userId, nil
}

func (controller *userController) GetById(c *gin.Context) {
	id, err := getUserId(c.Param(("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := controller.us.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
