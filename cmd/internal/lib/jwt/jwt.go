package jwt

import (
	"SSO/cmd/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	//формирования объекта в котором будут храниться все данные
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID                         //чтобы понимать кто логинится
	claims["email"] = user.Email                    //его имейл
	claims["exp"] = time.Now().Add(duration).Unix() //сколько будет существовать токен с сейчашнего времени
	claims["app_id"] = app.ID                       //апп Ид в который логинимся

	//подписываем токен
	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
