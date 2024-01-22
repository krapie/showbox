package handlers

import (
	"context"

	"github.com/krapie/showbox/server/database/mongo"
	"github.com/krapie/showbox/server/services/show"
	"github.com/krapie/showbox/server/services/ticket"
	lruCache "github.com/krapie/showbox/server/utils/cache"
)

type Handler struct {
	ctx       context.Context
	showSrv   *show.Service
	ticketSrv *ticket.Service
}

func New(ctx context.Context) (*Handler, error) {
	db, err := mongo.New(ctx)
	if err != nil {
		return nil, err
	}

	cache := lruCache.New()

	showSrv, err := show.New(ctx, db, cache)
	if err != nil {
		return nil, err
	}

	ticketSrv, err := ticket.New(ctx, db, cache)
	if err != nil {
		return nil, err
	}

	return &Handler{
		ctx:       ctx,
		showSrv:   showSrv,
		ticketSrv: ticketSrv,
	}, nil
}
