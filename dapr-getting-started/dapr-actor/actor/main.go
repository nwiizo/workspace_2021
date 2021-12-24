package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dapr/go-sdk/actor"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/pkg/errors"

	"github.com/kzmake/dapr-actor/domain"
)

type PiggyBankActor struct {
	actor.ServerImplBase
}

func NewPiggyBankActor() func() actor.Server {
	return func() actor.Server {
		return &PiggyBankActor{}
	}
}

func (a *PiggyBankActor) Type() string {
	return "PiggyBank"
}

func (a *PiggyBankActor) Put(ctx context.Context, coin domain.Coin) error {
	log.Println("Actor: ", a.Type(), "/", a.ID(), " call Put: ", coin)

	pg, err := a.get()
	if err != nil {
		return err
	}

	if pg.State == domain.Broken {
		log.Println("... broken piggy bank")
		return errors.Errorf("broken piggy bank")
	}

	pg.Coins = append(pg.Coins, coin)

	err = a.set(pg)
	if err != nil {
		return err
	}

	return nil
}
func (a *PiggyBankActor) Break(context.Context) ([]domain.Coin, error) {
	log.Println("Actor: ", a.Type(), "/", a.ID(), " call Break")

	pg, err := a.get()
	if err != nil {
		return nil, err
	}

	broken := &domain.PiggyBank{
		ID:    domain.PiggyBankID(a.ID()),
		State: domain.Broken,
		Coins: []domain.Coin{},
	}

	if err := a.set(broken); err != nil {
		return nil, err
	}

	return pg.Coins, nil
}
func (a *PiggyBankActor) Jingle(context.Context) (string, error) {
	log.Println("Actor: ", a.Type(), "/", a.ID(), " call Jingle")

	pg, err := a.get()
	if err != nil {
		return "", err
	}

	log.Println("Actor: ", pg, len(pg.Coins))

	if pg.State == domain.Broken {
		log.Println("... broken piggy bank")
		return "", errors.Errorf("broken piggy bank")
	}

	return strings.Repeat("じゃら", len(pg.Coins)), nil
}

func (a *PiggyBankActor) get() (*domain.PiggyBank, error) {
	pg := &domain.PiggyBank{
		ID:    domain.PiggyBankID(a.ID()),
		State: domain.Healthy,
		Coins: []domain.Coin{},
	}

	if found, err := a.GetStateManager().Contains("self"); err != nil {
		return nil, err
	} else if found {
		if err := a.GetStateManager().Get("self", pg); err != nil {
			return nil, err
		}
	}

	return pg, nil
}

func (a *PiggyBankActor) set(pg *domain.PiggyBank) error {
	if err := a.GetStateManager().Set("self", pg); err != nil {
		return err
	}

	return nil
}

func main() {
	s := daprd.NewService(":8080")
	s.RegisterActorImplFactory(NewPiggyBankActor())

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}
