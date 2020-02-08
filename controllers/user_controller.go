package controllers

import (
	"errors"
	"go_api_boilerplate/domain/user"
	"go_api_boilerplate/services/auth_service"
	"go_api_boilerplate/services/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserInput represents login/register request body format
type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserOutput represents returning user
type UserOutput struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Active    bool   `json:"active"`
}

// UserUpdateInput represents updating profile request body format
type UserUpdateInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// UserController interface
type UserController interface {
	Register(*gin.Context)
	Login(*gin.Context)
	GetByID(*gin.Context)
	GetProfile(*gin.Context)
	Update(*gin.Context)
	ForgotPassword(*gin.Context)
	ResetPassword(*gin.Context)
}

type userController struct {
	us user_service.UserService
	as auth_service.AuthService
}

// NewUserController instantiates User Controller
func NewUserController(us user_service.UserService, as auth_service.AuthService) UserController {
	return &userController{
		us: us,
		as: as,
	}
}

func (ctl *userController) Register(c *gin.Context) {
	// Read user input
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

	// Login
	err := ctl.login(c, &u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (ctl *userController) Login(c *gin.Context) {
	// Read user input
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from DB
	user, err := ctl.us.GetByEmail(userInput.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check password
	err = ctl.us.ComparePassword(userInput.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Login
	err = ctl.login(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	userOutput := ctl.mapToUserOutput(user)
	c.JSON(http.StatusOK, userOutput)
}

// @Summary Get user info of the logged in user
// @Produce  json
// @Success 200 {object} UserOutput
// @Failure 500 {object} ErrorResponse
// @Router /api/account/profile [get]
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
	userOutput := ctl.mapToUserOutput(user)
	c.JSON(http.StatusOK, userOutput)
}

func (ctl *userController) Update(c *gin.Context) {
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

	// Read user input
	var userInput UserUpdateInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.ID != id {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("Unauthorized")})
		return
	}

	user.FirstName = userInput.FirstName
	user.LastName = userInput.LastName
	user.Email = userInput.Email

	if err := ctl.us.Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userOutput := ctl.mapToUserOutput(user)
	c.JSON(http.StatusOK, userOutput)
}

func (ctl *userController) ForgotPassword(c *gin.Context) {}

func (ctl *userController) ResetPassword(c *gin.Context) {}

/*******************************/
//       PRIVATE METHODS
/*******************************/

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

func (ctl *userController) mapToUserOutput(u *user.User) *UserOutput {
	return &UserOutput{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
		Active:    u.Active,
	}
}

// Issue token and return user
func (ctl *userController) login(c *gin.Context, u *user.User) error {
	token, err := ctl.as.IssueToken(*u)
	if err != nil {
		return err
	}
	userOutput := ctl.mapToUserOutput(u)

	// Return user info + token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  userOutput,
	})
	return nil
}
