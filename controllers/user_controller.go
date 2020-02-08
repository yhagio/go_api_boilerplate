package controllers

import (
	"errors"
	"fmt"
	"go_api_boilerplate/domain/user"
	"go_api_boilerplate/services/auth_service"
	"go_api_boilerplate/services/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserOutput struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Active    bool   `json:"active"`
}

type UserController interface {
	Register(*gin.Context)
	Login(*gin.Context)
	GetByID(*gin.Context)
	GetProfile(*gin.Context)
	Update(*gin.Context)
}

type userController struct {
	us user_service.UserService
	as auth_service.AuthService
}

func NewUserController(us user_service.UserService, as auth_service.AuthService) UserController {
	return &userController{
		us: us,
		as: as,
	}
}

func (ctl *userController) Register(c *gin.Context) {
	// Validates (input, dupe user, email)
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := ctl.inputToUser(userInput)

	// Create user
	if err := ctl.us.Create(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Login (issue token)
	token, err := ctl.as.IssueToken(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userOutput := &UserOutput{
		ID: u.ID,
		// Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
		Active:    u.Active,
	}

	// Return user info + token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  userOutput,
	})
}

func (ctl *userController) Login(c *gin.Context) {
	// Validates (input, dupe user, email)
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := ctl.inputToUser(userInput)
	fmt.Println(u)
	// Get user
	// Login (issue token)
	// Return user info + token
}

func (ctl *userController) GetByID(c *gin.Context) {
	id, err := ctl.getUserID(c.Param(("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := ctl.us.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctl *userController) GetProfile(c *gin.Context) {
	id, exists := c.Get("user_id")
	if exists == false {
		c.JSON(http.StatusInternalServerError, errors.New("Invalid User ID"))
		return
	}

	user, err := ctl.us.GetByID(id.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctl *userController) Update(c *gin.Context) {
	// Validates (input, dupe user, email)
	var user UserInput
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Get user
	// Login (issue token)
	// Return user info + token
}

func (ctl *userController) getUserID(userIDParam string) (uint, error) {
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return 0, errors.New("user id should be a number")
	}
	return uint(userID), nil
}

func (ctl *userController) inputToUser(input UserInput) user.User {
	return user.User{
		Email:    input.Email,
		Password: input.Password,
	}
}
