package usermongorepo

import (
	"context"
	"time"

	"github.com/yhagio/go_api_boilerplate/domain/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	GetByEmail(email string) (*user.User, error)
}

type userRepo struct {
	db *mongo.Client
}

// NewUserRepo will instantiate User Repository
func NewUserMongoRepo(db *mongo.Client) Repo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetByEmail(email string) (*user.User, error) {
	var user user.User
	filter := bson.M{"email": email}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := u.db.Database("my-db").Collection("user").FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
