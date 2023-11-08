package db

import (
	"context"

	"github.com/offluck/ilove2rest/internal/entities/user"
)

type RedisClient struct{}

var _ Client = NewRedisClient()

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (*RedisClient) GetUsers(ctx context.Context) ([]user.UserDB, error) {
	panic("unimplemented")
}

func (*RedisClient) AddUser(ctx context.Context, user user.UserRequest) (user.UserDB, error) {
	panic("unimplemented")
}

func (*RedisClient) DeleteUser(ctx context.Context, username string) error {
	panic("unimplemented")
}

func (*RedisClient) GetUser(ctx context.Context, username string) (user.UserDB, error) {
	panic("unimplemented")
}

func (*RedisClient) UpdateUser(ctx context.Context, username string, newUser user.UserDB) (user.UserDB, error) {
	panic("unimplemented")
}
