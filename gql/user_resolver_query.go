package gql

import (
	"context"
	"errors"
	"github.com/yhagio/go_api_boilerplate/gql/gen"
)

func (r *queryResolver) User(ctx context.Context, id int) (*gen.User, error) {
	user, err := r.UserService.GetByID(uint(id))
	if err != nil {
		return nil, err
	}

	return &gen.User{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Active:    user.Active,
	}, nil
}

func (r *queryResolver) UserProfile(ctx context.Context) (*gen.User, error) {
	userID := ctx.Value("user_id")
	if userID == nil {
		return nil, errors.New("Unauthorized: Token is invlaid")
	}

	user, err := r.UserService.GetByID(userID.(uint))
	if err != nil {
		return nil, err
	}

	return &gen.User{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Active:    user.Active,
	}, nil
}
