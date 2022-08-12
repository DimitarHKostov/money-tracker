package jwt

import (
	"errors"
	"money-tracker/generator"
	"money-tracker/types"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

type JWTManager struct {
	PayloadGenerator generator.PayloadGeneratorInterface
	SecretKey        string
}

func (jwtm *JWTManager) GenerateToken(username string, duration time.Duration) (string, error) {
	payload, err := jwtm.PayloadGenerator.GeneratePayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(jwtm.SecretKey))
}

func (jwtm *JWTManager) VerifyToken(token string) (*types.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token is invalid")
		}
		return []byte(jwtm.SecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &types.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errors.New("token has expired")) {
			return nil, errors.New("token has expired")
		}
		return nil, errors.New("token is invalid")
	}

	payload, ok := jwtToken.Claims.(*types.Payload)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	return payload, nil
}
