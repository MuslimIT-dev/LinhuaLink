package utils

import (
	"LinhuaLink/backend/pkg/config"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 12 hours
	})
	return token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
}

func ParseAndVerifyJWT(tokenString string) (map[string]interface{}, error) {
	// keyFunc — функция, которая отдаёт ключ (секрет)
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		// проверяем алгоритм
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("unexpected signing method:", t.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.GetEnv("JWT_SECRET_ACCESS")), nil
	}

	// Парсим и проверяем
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		log.Println("error parsing token:", err)
		return nil, err
	}

	// Проверим валидность подписи
	if !token.Valid {
		log.Println("token invalid")
		return nil, errors.New("token invalid")
	}

	// Преобразуем claims в map
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		m := make(map[string]interface{})
		for k, v := range claims {
			m[k] = v
		}
		return m, nil
	}

	return nil, fmt.Errorf("cannot read claims")
}
