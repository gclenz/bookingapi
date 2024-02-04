package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IJWTHandler interface {
	GenerateJWT(userID string) (string, error)
	VerifyJWT(tokenString string) (string, error)
}

type JWTHandler struct {
	Secret string
}

func (jh *JWTHandler) GenerateJWT(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	signedToken, err := token.SignedString([]byte(jh.Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (jh *JWTHandler) VerifyJWT(tokenString string) (string, error) {
	var keyfunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", ErrInvalidToken
		}
		return []byte(jh.Secret), nil
	}

	token, err := jwt.Parse(tokenString, keyfunc)
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", ErrInvalidToken
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp := time.Unix(int64(claims["exp"].(float64)), 0)
		if exp.UTC().Before(time.Now().UTC()) {
			return "", ErrExpiredToken
		}
		return claims["userID"].(string), nil
	}
	return "", ErrCouldNotParseToken
}

var (
	ErrCouldNotParseToken = errors.New("could not parse token")
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("expired token")
)
