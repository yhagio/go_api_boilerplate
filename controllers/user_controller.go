package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/yhagio/go_api_boilerplate/domain/user"
	"github.com/yhagio/go_api_boilerplate/services/authservice"
	"github.com/yhagio/go_api_boilerplate/services/emailservice"
	"github.com/yhagio/go_api_boilerplate/services/userservice"

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
	us userservice.UserService
	as authservice.AuthService
	es emailservice.EmailService
}

// NewUserController instantiates User Controller
func NewUserController(
	us userservice.UserService,
	as authservice.AuthService,
	es emailservice.EmailService) UserController {
	return &userController{
		us: us,
		as: as,
		es: es,
	}
}

// @Summary Register new user
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/register [post]
func (ctl *userController) Register(c *gin.Context) {
	// Read user input
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	u := ctl.inputToUser(userInput)

	// Create user
	if err := ctl.us.Create(&u); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Send welcome email
	if err := ctl.es.Welcome(u.Email); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Login
	err := ctl.login(c, &u)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

// @Summary Login
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/login [post]
func (ctl *userController) Login(c *gin.Context) {
	// Read user input
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Get user from DB
	user, err := ctl.us.GetByEmail(userInput.Email)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Check password
	err = ctl.us.ComparePassword(userInput.Password, user.Password)
	if err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Login
	err = ctl.login(c, user)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

// @Summary Get user info of given id
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/users/{id} [get]
func (ctl *userController) GetByID(c *gin.Context) {
	id, err := ctl.getUserID(c.Param(("id")))
	if err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := ctl.us.GetByID(id)
	if err != nil {
		es := err.Error()
		if strings.Contains(es, "not found") {
			HTTPRes(c, http.StatusNotFound, es, nil)
			return
		}
		HTTPRes(c, http.StatusInternalServerError, es, nil)
		return
	}
	userOutput := ctl.mapToUserOutput(user)
	HTTPRes(c, http.StatusOK, "ok", userOutput)
}

// @Summary Get user info of the logged in user
// @Produce  json
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/account/profile [get]
func (ctl *userController) GetProfile(c *gin.Context) {
	id, exists := c.Get("user_id")
	if exists == false {
		HTTPRes(c, http.StatusBadRequest, "Invalid User ID", nil)
		return
	}

	user, err := ctl.us.GetByID(id.(uint))
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	userOutput := ctl.mapToUserOutput(user)
	HTTPRes(c, http.StatusOK, "ok", userOutput)
}

// @Summary Update account profile
// @Produce  json
// @Param email body string true "Email"
// @Param firstName body string false "First Name"
// @Param lastName body string false "Last Name"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/account/profile [put]
func (ctl *userController) Update(c *gin.Context) {
	// Get user id from context
	id, exists := c.Get("user_id")
	if exists == false {
		HTTPRes(c, http.StatusBadRequest, "Invalid User ID", nil)
		return
	}

	// Retrieve user given id
	user, err := ctl.us.GetByID(id.(uint))
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Read user input
	var userInput UserUpdateInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Check user
	if user.ID != id {
		HTTPRes(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	// Update user record
	user.FirstName = userInput.FirstName
	user.LastName = userInput.LastName
	user.Email = userInput.Email
	if err := ctl.us.Update(user); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Response
	userOutput := ctl.mapToUserOutput(user)
	HTTPRes(c, http.StatusOK, "ok", userOutput)
}

// @Summary Sends token to user's email to update user's password
// @Produce  json
// @Param email body string true "Email"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/forgot_password [post]
func (ctl *userController) ForgotPassword(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Issue token for user to update his/her password
	token, err := ctl.us.InitiateResetPassowrd(input.Email)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Send email with token to update password
	if err = ctl.es.ResetPassword(input.Email, token); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	HTTPRes(c, http.StatusOK, "Email sent", nil)
	return
}

// @Summary Update user's password
// @Produce  json
// @Param password body string true "Password"
// @Param token query string true "Token"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/update_password [post]
func (ctl *userController) ResetPassword(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	token := c.Request.URL.Query().Get("token")
	if token == "" {
		HTTPRes(c, http.StatusNotFound, "Requires token", nil)
		return
	}

	user, err := ctl.us.CompleteUpdatePassword(token, input.Password)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = ctl.login(c, user)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

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
	out := gin.H{"token": token, "user": userOutput}
	HTTPRes(c, http.StatusOK, "ok", out)
	return nil
}
