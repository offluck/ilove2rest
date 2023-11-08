package db

import (
	"context"

	"github.com/offluck/ilove2rest/internal/entities/user"
)

type Client interface {
	GetUsers(ctx context.Context) ([]user.UserDB, error)
	GetUser(ctx context.Context, username string) (user.UserDB, error)
	AddUser(ctx context.Context, user user.UserRequest) (user.UserDB, error)
	UpdateUser(ctx context.Context, username string, newUser user.UserDB) (user.UserDB, error)
	DeleteUser(ctx context.Context, username string) error
}
