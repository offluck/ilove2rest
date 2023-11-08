package db

import (
	"context"

	"github.com/offluck/ilove2rest/internal/entities"
)

type Client interface {
	GetUser(ctx context.Context, username string)
	AddUser(ctx context.Context, user entities.User)
	UpdateUser(ctx context.Context, username string, newUser entities.User)
	DeleteUser(ctx context.Context, username string)
}
