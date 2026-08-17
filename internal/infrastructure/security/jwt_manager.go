package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JWTManager struct {
	secretKey []byte
	accessTtl time.Duration
	refreshTtl time.Duration
}

func NewJWTManager(secret []byte, accessTtl, refreshTtl time.Duration) *JWTManager {
	return &JWTManager{
		secretKey: secret,
		accessTtl: accessTtl,
		refreshTtl: refreshTtl,
	}
}

func (manager *JWTManager) NewAccessToken(userID int) (string, error) {
	return manager.newToken(userID, manager.accessTtl)
}

func (manager *JWTManager) NewRefreshToken(userID int) (string, error) {
	return manager.newToken(userID, manager.refreshTtl)
}

func (manager *JWTManager) newToken(userID int, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(manager.secretKey)
}

func (manager *JWTManager) ParseToken(inputToken string) (int, error) {
	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return manager.secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	payload := token.Claims.(jwt.MapClaims)

	return payload["sub"].(int), nil
}