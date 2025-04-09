package models

type App struct {
	ID     int
	Name   string
	Secret string //подписать токен и потом валидировать его
}
