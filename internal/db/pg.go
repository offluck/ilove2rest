package db

import (
	"context"

	"github.com/offluck/ilove2rest/internal/entities/user"
)

type PGClient struct{}

var _ Client = NewPGClient()

func NewPGClient() *PGClient {
	return &PGClient{}
}

func (*PGClient) GetUsers(ctx context.Context) ([]user.UserDB, error) {
	panic("unimplemented")
}

func (*PGClient) AddUser(ctx context.Context, user user.UserRequest) (user.UserDB, error) {
	panic("unimplemented")
}

func (*PGClient) DeleteUser(ctx context.Context, username string) error {
	panic("unimplemented")
}

func (*PGClient) GetUser(ctx context.Context, username string) (user.UserDB, error) {
	panic("unimplemented")
}

func (*PGClient) UpdateUser(ctx context.Context, username string, newUser user.UserDB) (user.UserDB, error) {
	panic("unimplemented")
}
