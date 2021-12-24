package api

import (
	"context"

	"github.com/kzmake/dapr-actor/domain"
)

type PiggyBankActor struct {
	Id string `json:"id"`

	Put    func(context.Context, domain.Coin) error
	Break  func(context.Context) ([]domain.Coin, error)
	Jingle func(context.Context) (string, error)
}

func NewPiggyBankActor(pb *domain.PiggyBank) *PiggyBankActor {
	return &PiggyBankActor{Id: string(pb.ID)}
}

func (a *PiggyBankActor) Type() string {
	return "PiggyBank"
}

func (a *PiggyBankActor) ID() string {
	return a.Id
}
