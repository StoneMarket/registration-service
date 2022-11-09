package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/StoneMarket/registration-service/internal/models"
	"github.com/jackc/pgx/v4"
)

const (
	findUserLogin = `SELECT login FROM users WHERE login = $1`
	createUser    = `INSERT INTO users (firstname, login, password) VALUES ($1, $2, $3)`
	findUserData  = `SELECT * FROM users WHERE login = $1`
)

type Storage struct {
	pgdb *pgx.Conn
}

func Connect(ctx context.Context, cfg *pgx.ConnConfig) (*pgx.Conn, error) {
	conn, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewStorage(pgdb *pgx.Conn) *Storage {
	return &Storage{
		pgdb: pgdb,
	}
}

func (s *Storage) SetUserData(ctx context.Context, user *models.User) error {

	if err := s.pgdb.QueryRow(ctx, findUserLogin, user.Login).Scan(); err != nil {
		if err == pgx.ErrNoRows {

			_, err = s.pgdb.Exec(ctx, createUser, user.Name, user.Login, user.Password)
			if err != nil {
				return fmt.Errorf("storage: method: SetUserData: creating, error: %s", err)
			}
			return nil
		}

		return fmt.Errorf("storage: method: SetUserData: finding, error: %s", err)
	}

	return nil
}

func (s *Storage) FindUserData(ctx context.Context, login, password string) (*models.User, error) {
	row := s.pgdb.QueryRow(ctx, findUserData, login)
	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Login, &user.Password); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
