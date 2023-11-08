package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/offluck/ilove2rest/internal/entities/user"
	"go.uber.org/zap"
)

type PGClient struct {
	*sql.DB
	logger *zap.Logger
}

var _ Client = NewPGClient(nil, nil)

func NewPGClient(db *sql.DB, logger *zap.Logger) *PGClient {
	return &PGClient{
		DB:     db,
		logger: logger,
	}
}

func (client *PGClient) GetUsers(ctx context.Context) ([]user.UserDB, error) {
	rows, err := client.QueryContext(ctx, "SELECT username, password, first_name, last_name, email, phone FROM users")
	if err != nil {
		return nil, fmt.Errorf("Failed to create users query: %+v", err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			client.logger.Error("failed to close connection", zap.Error(err))
		}
	}()

	usersDB := make([]user.UserDB, 0)
	userDB := user.UserDB{}
	for rows.Next() {
		err := rows.Scan(
			&userDB.Username,
			&userDB.Password,
			&userDB.FirstName,
			&userDB.LastName,
			&userDB.Email,
			&userDB.Phone,
		)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan row: %+v", err)
		}
		usersDB = append(usersDB, userDB)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("Rows error: %+v", err)
	}

	return usersDB, nil
}

func (client *PGClient) GetUser(ctx context.Context, username string) (user.UserDB, error) {
	query := client.QueryRowContext(ctx, "SELECT username, password, first_name, last_name, email, phone FROM users WHERE username LIKE ?", username)
	if query.Err() != nil {
		return user.UserDB{}, fmt.Errorf("Failed to create user query: %+v", query.Err())
	}

	userDB := user.UserDB{}
	err := query.Scan(
		&userDB.Username,
		&userDB.Password,
		&userDB.FirstName,
		&userDB.LastName,
		&userDB.Email,
		&userDB.Phone,
	)
	if err != nil {
		return user.UserDB{}, fmt.Errorf("Failed to query a user: %+v", err)
	}

	return userDB, nil
}

func (client *PGClient) AddUser(ctx context.Context, userDB user.UserDB) (user.UserDB, error) {
	query := `INSERT INTO users (username, password, first_name, last_name, email, phone)
				VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := client.ExecContext(
		ctx,
		query,
		userDB.Username,
		userDB.Password,
		userDB.FirstName,
		userDB.LastName,
		userDB.Email,
		userDB.Phone,
	)
	if err != nil {
		return user.UserDB{}, fmt.Errorf("Failed to insert user: %+v", err)
	}

	return userDB, nil
}

func (client *PGClient) UpdateUser(ctx context.Context, username string, newUserDB user.UserDB) (user.UserDB, error) {
	query := `UPDATE users
				SET username = $1,
					password = $2,
					first_name = $3,
					last_name = $4,
					email = $5,
					phone = $6,
				WHERE username LIKE $7`

	_, err := client.ExecContext(
		ctx,
		query,
		newUserDB.Username,
		newUserDB.Password,
		newUserDB.FirstName,
		newUserDB.LastName,
		newUserDB.Email,
		newUserDB.Phone,
	)
	if err != nil {
		return user.UserDB{}, fmt.Errorf("Failed to update user: %+v", err)
	}
	return newUserDB, nil
}

func (client *PGClient) DeleteUser(ctx context.Context, username string) error {
	query := `DELETE FROM users WHERE username LIKE ?)`

	_, err := client.ExecContext(
		ctx,
		query,
		username,
	)
	if err != nil {
		return fmt.Errorf("Failed to delete user: %+v", err)
	}

	return nil
}
