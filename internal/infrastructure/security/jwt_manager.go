package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey  []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWTManager(secret []byte, accessTTL, refreshTTL time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:  secret,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

func (manager *JWTManager) NewAccessToken(userID int) (string, error) {
	return manager.newToken(userID, manager.accessTTL)
}

func (manager *JWTManager) NewRefreshToken(userID int) (string, error) {
	return manager.newToken(userID, manager.refreshTTL)
}

func (manager *JWTManager) newToken(userID int, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(manager.secretKey)
}

func (m *JWTManager) ParseToken(inputToken string) (int, error) {
	token, err := jwt.Parse(inputToken, func(t *jwt.Token) (any, error) {
		return m.secretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return 0, fmt.Errorf("parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	raw, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user_id claim")
	}
	return int(raw), nil
}
