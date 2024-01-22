package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/krapie/showbox/server/models"
)

type Client struct {
	client *mongo.Client
}

func New(ctx context.Context) (*Client, error) {
	uri := "mongodb://localhost:27017"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) FindShowById(ctx context.Context, id primitive.ObjectID) (*models.Show, error) {
	var show models.Show

	err := c.client.Database("showbox").Collection("shows").FindOne(ctx, bson.M{"_id": id}).Decode(&show)
	if err != nil {
		return nil, err
	}

	return &show, nil
}

func (c *Client) CreateShow(ctx context.Context, show *models.Show) (*models.Show, error) {
	res, err := c.client.Database("showbox").Collection("shows").InsertOne(ctx, show)
	if err != nil {
		return nil, err
	}
	show.ID = res.InsertedID.(primitive.ObjectID)

	return show, nil
}

func (c *Client) FindTicketById(ctx context.Context, id primitive.ObjectID) (*models.Ticket, error) {
	var ticket models.Ticket

	err := c.client.Database("showbox").Collection("tickets").FindOne(ctx, bson.M{"_id": id}).Decode(&ticket)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (c *Client) FindTicketsByShowId(ctx context.Context, id primitive.ObjectID) ([]*models.Ticket, error) {
	var tickets []*models.Ticket

	cursor, err := c.client.Database("showbox").Collection("tickets").Find(ctx, bson.M{"show_id": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var ticket models.Ticket
		err = cursor.Decode(&ticket)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, &ticket)
	}

	return tickets, nil
}

func (c *Client) CreateTicket(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res, err := c.client.Database("showbox").Collection("tickets").InsertOne(ctx, ticket)
	if err != nil {
		return nil, err
	}
	ticket.ID = res.InsertedID.(primitive.ObjectID)

	return ticket, nil
}

func (c *Client) UpdateTicket(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	_, err := c.client.Database("showbox").Collection("tickets").UpdateOne(ctx, bson.M{"_id": ticket.ID}, bson.M{"$set": ticket})
	if err != nil {
		return nil, err
	}

	return ticket, nil
}
