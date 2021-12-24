package domain

import (
	crand "crypto/rand"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

type PiggyBank struct {
	ID    PiggyBankID    `json:"id"`
	State PiggyBankState `json:"state"`
	Coins []Coin         `json:"coins"`
}

func NewPiggyBank() *PiggyBank {
	return &PiggyBank{
		ID:    PiggyBankID(ulid.MustNew(ulid.Timestamp(time.Now()), crand.Reader).String()),
		State: Healthy,
		Coins: []Coin{},
	}
}
