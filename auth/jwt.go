package auth

import (
"github.com/dgrijalva/jwt-go"

"crypto/rsa"
"errors"
"fmt"
"time"
)

// todo: key rotation
type KeysRSA struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey

	expire time.Duration
}

type Claims struct {
	Id string
	jwt.StandardClaims
}

func NewKeysRSA(privateBytes, publicBytes []byte, keyExpiration time.Duration) (*KeysRSA, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return nil, err
	}
	return &KeysRSA{
		publicKey:  publicKey,
		privateKey: privateKey,
		expire:     keyExpiration,
	}, nil
}

func (j KeysRSA) IssueToken(userId string) (string, error) {
	claims := Claims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.expire).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(j.privateKey)
}

func (j KeysRSA) UserIdByToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected token signing method")
		}
		return j.publicKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	return claims.Id, nil
}
