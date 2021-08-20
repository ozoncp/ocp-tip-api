package db

import (
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// Connect создает и возвращает пул соединений с БД по строке подключения
func Connect(DSN string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", DSN)
	if err != nil {
		return nil, err
	}
	return db, nil
}
