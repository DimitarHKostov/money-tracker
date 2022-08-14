package generator

import (
	"money-tracker/types"
	"time"
)

type PayloadGeneratorInterface interface {
	GeneratePayload(username string, duration time.Duration) (*types.Payload, error)
}
