package sqlite

import (
	"SSO/cmd/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

// Конструктор Storage
func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	//Указываем путь до файла БД
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

// Реализация метода saveUser
func (s *Storage) SaveUser(ctx context.Context, email string, passsHash []byte) (int64, error) {
	const op = "storage.sqlite.SaveUser"

	//Простенький запрос на добавление пользователя
	stmt, err := s.db.Prepare("INSERT INTO users(email, pass hash) VALUES(?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	//Выполнение запроса
	res, err := stmt.ExecContext(ctx, email, passsHash)
	if err != nil {
		var sqliteErr sqlite3.Error
		var ErrUserExists = errors.New("user already exists")
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	//Получаем ID созданного пользователя
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

// User returns user by email
func (s *Storage) User(ctx context.Context, email string) (models.User, error) {

}
