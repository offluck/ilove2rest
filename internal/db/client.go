package db

import (
	"context"

	"github.com/offluck/ilove2rest/internal/entities/user"
)

type Client interface {
	GetUser(ctx context.Context, username string) (user.UserResponse, error)
	AddUser(ctx context.Context, user user.UserRequest) (user.UserResponse, error)
	UpdateUser(ctx context.Context, username string, newUser user.UserRequest) (user.UserResponse, error)
	DeleteUser(ctx context.Context, username string) error
}
