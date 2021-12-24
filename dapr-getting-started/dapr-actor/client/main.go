package main

import (
	"context"
	"log"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	"golang.org/x/sync/errgroup"

	"github.com/kzmake/dapr-actor/api"
	"github.com/kzmake/dapr-actor/domain"
)

func main() {
	time.Sleep(15 * time.Second) // wait dapr-sidecar

	c, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		pb := domain.NewPiggyBank()
		actor := api.NewPiggyBankActor(pb)
		c.ImplActorClientStub(actor)

		s, err := actor.Jingle(ctx)
		log.Printf("actor(pb: %s).Jingle: %+v, %+v", string(pb.ID), s, err)
		if err != nil {
			return err
		}

		actor.Put(ctx, domain.Yen10)
		s, err = actor.Jingle(ctx)
		log.Printf("actor(pb: %s).Jingle: %+v, %+v", string(pb.ID), s, err)
		if err != nil {
			return err
		}

		actor.Put(ctx, domain.Yen100)
		actor.Put(ctx, domain.Yen500)
		s, err = actor.Jingle(ctx)
		log.Printf("actor(pb: %s).Jingle: %+v, %+v", string(pb.ID), s, err)
		if err != nil {
			return err
		}

		coins, err := actor.Break(ctx)
		log.Printf("actor(pb: %s).Break: %+v, %+v", string(pb.ID), coins, err)
		if err != nil {
			return err
		}

		return nil
	})

	eg.Go(func() error {
		pb := domain.NewPiggyBank()
		actor := api.NewPiggyBankActor(pb)
		c.ImplActorClientStub(actor)

		coins, err := actor.Break(ctx)
		log.Printf("actor(pb: %s).Break: %+v, %+v", string(pb.ID), coins, err)
		if err != nil {
			return err
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
