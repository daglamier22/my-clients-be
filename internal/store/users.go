package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Uuid      string `json:"uuid"`
}

type UsersStore struct {
	db *sql.DB
}

func NewUsersStore(db *sql.DB) UsersStore {
	return UsersStore{
		db: db,
	}
}

func (s *UsersStore) Create(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users (username, password, first_name, last_name, email)
		VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at;
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		u.Username,
		u.Password,
		u.FirstName,
		u.LastName,
		u.Email,
	).Scan(
		&u.Id,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("User already exists")
		}
		return err
	}

	return nil
}

func (s *UsersStore) GetUser(ctx context.Context, u *User) error {
	query := `
		SELECT user, password, first_name, last_name, email
		FROM users
		WHERE id=$1;
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		u.Id,
	).Scan(
		&u.Username,
		&u.Password,
		&u.FirstName,
		&u.LastName,
		&u.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No rows found with ID %d", u.Id)
			return err
		}
		log.Fatal(err)
		return err
	}

	return nil
}

func (s *UsersStore) GetAllUsers(ctx context.Context, u *[]User) error {
	query := `
		SELECT id, username, password, first_name, last_name, email
		FROM users;
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		newUser := User{}
		err := rows.Scan(
			&newUser.Id,
			&newUser.Username,
			&newUser.Password,
			&newUser.FirstName,
			&newUser.LastName,
			&newUser.Email,
		)
		if err != nil {
			log.Fatal(err)
		}
		*u = append(*u, newUser)
	}

	return nil
}
