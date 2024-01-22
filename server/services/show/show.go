package show

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/krapie/showbox/server/database/mongo"
	"github.com/krapie/showbox/server/models"
	"github.com/krapie/showbox/server/utils/cache"
)

type Service struct {
	ctx   context.Context
	db    *mongo.Client
	cache *cache.Cache
}

func New(
	ctx context.Context,
	db *mongo.Client,
	cache *cache.Cache,
) (*Service, error) {
	return &Service{
		ctx:   ctx,
		db:    db,
		cache: cache,
	}, nil
}

func (s *Service) GetShow(ctx context.Context, id primitive.ObjectID) (*models.Show, error) {
	show, err := s.db.FindShowById(ctx, id)
	if err != nil {
		return nil, err
	}

	return show, nil
}

func (s *Service) RegisterShow(ctx context.Context, show *models.Show) error {
	// create show
	show, err := s.db.CreateShow(ctx, show)
	if err != nil {
		return err
	}

	// create 100 tickets for this show
	// TODO(krapie): create API for this action
	for i := 0; i < 100; i++ {
		ticket := models.NewTicket(primitive.NilObjectID, show.ID, primitive.NilObjectID)
		if _, err = s.db.CreateTicket(ctx, ticket); err != nil {
			return err
		}

		// add ticket to cache
		s.cache.Add(ticket.ID, ticket)
	}

	return nil
}
