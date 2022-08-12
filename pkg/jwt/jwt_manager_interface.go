package jwt

import (
	"money-tracker/types"
	"time"
)

type JWTManagerInterface interface {
	GenerateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*types.Payload, error)
}
