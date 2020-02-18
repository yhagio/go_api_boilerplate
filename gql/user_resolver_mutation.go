package gql

import (
	"context"
	"errors"

	"github.com/yhagio/go_api_boilerplate/domain/user"
	"github.com/yhagio/go_api_boilerplate/gql/gen"
)

func (r *mutationResolver) Login(ctx context.Context, input gen.RegisterLogin) (*gen.RegisterLoginOutput, error) {
	user, err := r.UserService.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	err = r.UserService.ComparePassword(input.Password, user.Password)
	if err != nil {
		return nil, err
	}

	token, err := r.AuthService.IssueToken(*user)
	if err != nil {
		return nil, err
	}

	return &gen.RegisterLoginOutput{
		Token: token,
		User: &gen.User{
			ID:        int(user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			Active:    user.Active,
		},
	}, nil
}

func (r *mutationResolver) Register(ctx context.Context, input gen.RegisterLogin) (*gen.RegisterLoginOutput, error) {
	userDomain := &user.User{
		Email:    input.Email,
		Password: input.Password,
	}

	err := r.UserService.Create(userDomain)
	if err != nil {
		return nil, err
	}

	token, err := r.AuthService.IssueToken(*userDomain)
	if err != nil {
		return nil, err
	}

	return &gen.RegisterLoginOutput{
		Token: token,
		User: &gen.User{
			ID:        int(userDomain.ID),
			FirstName: userDomain.FirstName,
			LastName:  userDomain.LastName,
			Email:     userDomain.Email,
			Role:      userDomain.Role,
			Active:    userDomain.Active,
		},
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input gen.UpdateUser) (*gen.User, error) {
	userID := ctx.Value("user_id")
	if userID == nil {
		return nil, errors.New("Unauthorized: Token is invlaid")
	}

	usr, err := r.UserService.GetByID(userID.(uint))
	if err != nil {
		return nil, err
	}

	if input.Email != "" {
		usr.Email = input.Email
	}
	if input.FirstName != nil {
		usr.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		usr.LastName = *input.LastName
	}

	err = r.UserService.Update(usr)
	if err != nil {
		return nil, err
	}

	return &gen.User{
		ID:        int(usr.ID),
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		Role:      usr.Role,
		Active:    usr.Active,
	}, nil
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (bool, error) {
	if email == "" {
		return false, errors.New("Email is required")
	}

	// Issue token for user to update his/her password
	token, err := r.UserService.InitiateResetPassowrd(email)
	if err != nil {
		return false, err
	}

	// Send email with token to update password
	if err = r.EmailService.ResetPassword(email, token); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) ResetPassword(ctx context.Context, resetToken string, password string) (*gen.RegisterLoginOutput, error) {
	if resetToken == "" {
		return nil, errors.New("Token is required")
	}

	if password == "" {
		return nil, errors.New("New password is required")
	}


	user, err := r.UserService.CompleteUpdatePassword(resetToken, password)
	if err != nil {
		return nil, err
	}

	token, err := r.AuthService.IssueToken(*user)
	if err != nil {
		return nil, err
	}

	return &gen.RegisterLoginOutput{
		Token: token,
		User: &gen.User{
			ID:        int(user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			Active:    user.Active,
		},
	}, nil
}
