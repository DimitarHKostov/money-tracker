package generator

import (
	"money-tracker/types"
	"time"
)

type PayloadGeneratorInterface interface {
	GeneratePayload(account *types.Account, duration time.Duration) (*types.Payload, error)
}
