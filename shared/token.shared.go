package shared

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const userIdKey = "user_id"
const userKey = "user"
const userPowersKey = "user_powers"

func UserIdKey() string {
	return userIdKey
}

func UserKey() string {
	return userKey
}

func UserPowersKeys() string {
	return userPowersKey
}

func GenerateToken(userID string, secret string, expirationTime int) (string, error) {
	claims := jwt.MapClaims{
		userIdKey: userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(expirationTime)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString string, secret string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		jwt.Keyfunc(func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		}),
	)
	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid token")
	}

	// aseguramos que los claims se puedan usar como MapClaims
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, fmt.Errorf("invalid claims format")
	}

	return token, mapClaims, nil
}
