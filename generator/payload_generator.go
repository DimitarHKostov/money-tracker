package generator

import (
	"money-tracker/types"
	"time"

	"github.com/google/uuid"
)

type PayloadGenerator struct{}

func (pg *PayloadGenerator) GeneratePayload(username string, duration time.Duration) (*types.Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &types.Payload{
		Id:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
