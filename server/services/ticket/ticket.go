package ticket

import (
	"context"
	"errors"

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

func (s *Service) GetAvailableTickets(ctx context.Context, showID primitive.ObjectID) ([]*models.Ticket, error) {
	// get all tickets
	tickets, err := s.db.FindTicketsByShowId(ctx, showID)
	if err != nil {
		return nil, err
	}

	// exclude tickets with userID or exists on cache
	var availableTickets []*models.Ticket
	for _, ticket := range tickets {
		if ticket.UserID.IsZero() && !s.cache.Exists(ticket.ID) {
			availableTickets = append(availableTickets, ticket)
		}
	}

	return availableTickets, nil
}

func (s *Service) ReserveTicket(ticketID, userID primitive.ObjectID) error {
	// TODO(krapie): this assumes that user only reserve for available tickets from `GetAvailableTickets()`
	// check if ticket is in cache or already reserved by other user
	t, ok := s.cache.Get(ticketID)
	if ok {
		ticket := t.(*models.Ticket)

		if ticket.UserID != userID {
			return errors.New("ticket is already reserved by other user")
		}
	}

	// if not, get create ticket and add to cache
	ticket := models.NewTicket(ticketID, primitive.NilObjectID, userID)
	s.cache.Add(ticketID, ticket)

	return nil
}

func (s *Service) BuyTicket(ctx context.Context, ticketID, userID primitive.ObjectID) error {
	// check if ticket is cache & userID is same with ticket's userID
	t, ok := s.cache.Get(ticketID)
	if !ok {
		return errors.New("ticket is not available")
	}
	ticket := t.(*models.Ticket)

	if ticket.UserID != userID {
		return errors.New("ticket is already reserved by other user")
	}

	// if correct, update ticket and update in db
	// and remove from cache
	_, err := s.db.UpdateTicket(ctx, ticket)
	if err != nil {
		return err
	}
	s.cache.Remove(ticketID)

	return nil
}
