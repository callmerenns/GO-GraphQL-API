package service

import (
	"fmt"
	"time"

	"github.com/altsaqif/go-graphql/cmd/entity"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type TokenConfig struct {
	IssuerName       string
	JwtSignatureKey  []byte
	JwtSigningMethod *jwt.SigningMethodHMAC
	JwtExpiresTime   time.Duration
}

type JwtService struct {
	config TokenConfig
}

func NewJwtService(config TokenConfig) *JwtService {
	return &JwtService{config: config}
}

func (s *JwtService) CreateToken(user *entity.User) (string, error) {
	claims := Claims{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.JwtExpiresTime)),
			Issuer:    s.config.IssuerName,
		},
	}
	token := jwt.NewWithClaims(s.config.JwtSigningMethod, claims)
	return token.SignedString(s.config.JwtSignatureKey)
}

func (s *JwtService) ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.config.JwtSignatureKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
