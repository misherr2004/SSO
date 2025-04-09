package sl // Объявление пакета

import (
	"log/slog" // Импорт стандартного пакета для логирования
)

// Err — функция, которая преобразует ошибку в slog.Attr
func Err(err error) slog.Attr {
	if err == nil {
		return slog.String("error", "nil") // Обработка nil-ошибок
	}
	return slog.String("error", err.Error())
}
