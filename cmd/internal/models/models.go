package models

// User представляет модель пользователя
type User struct {
	ID       int64
	Email    string
	PassHash []byte
}
