package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/daglamier22/my-clients-be/internal/application"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(dbConfig application.DbConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbConfig.Addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxIdleTime(dbConfig.MaxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, err
}

func NewTest() *sql.DB {
	return nil
}
